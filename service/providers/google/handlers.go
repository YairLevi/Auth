package google

import (
	"auth/database"
	"auth/database/types"
	"auth/service/providers"
	"auth/service/session"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"io"
	"net/http"
	"time"
)

var db = database.DB

func LoginHandler(ctx echo.Context) error {
	appID := ctx.Param("appId")
	googleOauthConfig, err := providers.GetConfigByAppID(appID, providers.Google)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	url := googleOauthConfig.AuthCodeURL(appID)
	return ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func CallbackHandler(ctx echo.Context) error {
	appID := ctx.QueryParam("state")
	if appID == "" {
		return ctx.JSON(http.StatusBadRequest, "invalid app ID")
	}

	googleOauthConfig, err := providers.GetConfigByAppID(appID, providers.Google)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	// Parse the authorization code from the query parameters
	code := ctx.QueryParam("code")

	// Exchange the authorization code for an access token
	token, err := googleOauthConfig.Exchange(ctx.Request().Context(), code)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	// Use the token to make a request to the Google API to get user details
	userInfoResp, err := googleOauthConfig.Client(ctx.Request().Context(), token).Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	defer userInfoResp.Body.Close()

	userInfoBytes, err := io.ReadAll(userInfoResp.Body)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	var userInfo map[string]interface{}
	if err = json.Unmarshal(userInfoBytes, &userInfo); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	fmt.Println("user:", userInfo)

	var user types.User
	err = db.Where("email = ?", userInfo["email"].(string)).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user.Email = userInfo["email"].(string)
		user.LastLogin = time.Now()
		user.FirstName = userInfo["given_name"].(string)
		user.LastName = userInfo["family_name"].(string)
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
