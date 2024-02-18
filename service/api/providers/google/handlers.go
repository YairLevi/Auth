package google

import (
	"auth/service/api/providers"
	"auth/service/database"
	"auth/service/database/types"
	"auth/service/session"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"io"
	"net/http"
	"time"
)

var db = database.DB

func LoginHandler(ctx echo.Context) error {
	googleOauthConfig, err := providers.GetOAuthConfig(providers.Google)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	url := googleOauthConfig.AuthCodeURL("state")
	return ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func CallbackHandler(ctx echo.Context) error {
	googleOauthConfig, err := providers.GetOAuthConfig(providers.Google)
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

	var user types.User
	err = db.Where("email = ?", userInfo["email"].(string)).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user.Email = userInfo["email"].(string)
		user.LastLogin = time.Now()
		user.Username = userInfo["given_name"].(string) + " " + userInfo["family_name"].(string)
		if err := db.Create(&user).Error; err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	} else if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	var security types.SecurityConfig
	if err := db.First(&security).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, "server error signing")
	}
	jwtCookie, err := session.GenerateJWT(session.Config{
		SigningKey: security.SessionKey,
		Expiration: 24 * time.Hour,
		Payload: session.Payload{
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
		Secure:   true,
	})
	return ctx.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
}
