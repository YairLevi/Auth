package github

import (
	"auth/database"
	"auth/database/types"
	auth "auth/service/middleware"
	"auth/service/providers"
	"auth/service/session"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"time"
)

var db = database.DB

func LoginHandler(ctx echo.Context) error {
	appID := ctx.(auth.Context).AppID
	githubOauthConfig, err := providers.GetConfigByAppID(appID, providers.Github)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid app ID")
	}
	url := githubOauthConfig.AuthCodeURL(appID)
	return ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func CallbackHandler(ctx echo.Context) error {
	appID := ctx.QueryParam("state")
	if appID == "" {
		return ctx.JSON(http.StatusBadRequest, "invalid app ID")
	}

	oauthConfig, err := providers.GetConfigByAppID(appID, providers.Github)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	// Parse the authorization code from the query parameters
	code := ctx.QueryParam("code")

	// Exchange the authorization code for an access token
	token, err := oauthConfig.Exchange(ctx.Request().Context(), code)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}

	if token.Extra("email") == nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}

	emailsResponse, err := oauthConfig.Client(ctx.Request().Context(), token).Get("https://api.github.com/user/emails")
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}
	defer emailsResponse.Body.Close()

	var emails []map[string]interface{}
	err = json.NewDecoder(emailsResponse.Body).Decode(&emails)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	primaryEmail := ""
	for _, m := range emails {
		if m["primary"] == true && m["verified"] == true {
			primaryEmail = m["email"].(string)
		}
	}
	if primaryEmail == "" {
		return ctx.JSON(http.StatusUnauthorized, "Failed to get a valid email address")
	}

	// Use the token to make a request to the Google API to get user details
	userInfoResp, err := oauthConfig.Client(ctx.Request().Context(), token).Get("https://api.github.com/user")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Failed to get user details from Google API")
	}
	defer userInfoResp.Body.Close()

	var userInfo map[string]interface{}
	err = json.NewDecoder(userInfoResp.Body).Decode(&userInfo)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	var user types.User
	err = db.Where("email = ?", primaryEmail).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user.Email = primaryEmail
		user.LastLogin = time.Now()
		userName, ok := userInfo["name"].(string)
		if !ok {
			userName = userInfo["login"].(string) // guaranteed to work. GitHub forces to have a login name.
		}
		user.FirstName = userName
		user.AppID = appID
		if err := db.Create(&user).Error; err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)

		}
	} else if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)

	}
	jwtCookie, err := session.GenerateJWT(session.Config{
		SigningKey: session.SecretKey,
		Expiration: 24 * time.Hour,
		Payload: session.Payload{
			AppID:  appID,
			UserID: user.ID,
		},
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)

	}
	ctx.SetCookie(&http.Cookie{
		Name:     session.CookieName,
		Value:    jwtCookie,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	})
	return ctx.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
}
