package standard

import (
	"auth/service/database/types"
	auth "auth/service/middleware"
	"auth/service/session"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func RegisterHandler(ctx echo.Context) error {
	registerDTO := struct {
		types.User
		Password string `json:"password"`
	}{}
	if err := ctx.Bind(&registerDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	appID := ctx.(auth.Context).AppID

	hashedPasswordInBytes := sha256.Sum256([]byte(registerDTO.Password))
	hashedPasswordEncodedString := base64.StdEncoding.EncodeToString(hashedPasswordInBytes[:])
	user := registerDTO.User
	user.PasswordHash = hashedPasswordEncodedString
	user.AppID = appID

	if err := db.Create(&user).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)

	}

	return ctx.JSON(http.StatusOK, &user)
}

func EmailPasswordLoginHandler(ctx echo.Context) error {
	loginDTO := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	if err := ctx.Bind(&loginDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	appID := ctx.(auth.Context).AppID

	appGuard := lockoutManager.AppGuards[appID]
	if appGuard != nil {
		if appGuard.IsLocked(loginDTO.Email) {
			return ctx.JSON(
				http.StatusForbidden,
				fmt.Sprintf(
					"this email is in lockout state. it will be released in %.1f seconds.",
					appGuard.Duration.Seconds(),
				),
			)
		}
	} else {
		fmt.Println("lockout guard is nil")
	}

	var user types.User
	res := db.Where("email = ? AND app_id = ?", loginDTO.Email, appID).First(&user)
	if res.Error != nil {
		appGuard.Fail(loginDTO.Email)
		return ctx.JSON(http.StatusUnauthorized, "unauthorized")
	}

	hashedPasswordInBytes := sha256.Sum256([]byte(loginDTO.Password))
	hashedPasswordEncodedString := base64.StdEncoding.EncodeToString(hashedPasswordInBytes[:])

	if hashedPasswordEncodedString != user.PasswordHash {
		appGuard.Fail(loginDTO.Email)
		return ctx.JSON(http.StatusUnauthorized, "unauthorized")
	}
	var app types.App
	if err := db.Where("id = ?", appID).Find(&app).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, "server error signing")
	}
	jwtCookie, err := session.GenerateJWT(session.Config{
		Expiration: time.Hour * 24,
		SigningKey: app.SessionKey,
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
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		Secure:   true,
	})

	return ctx.JSON(http.StatusOK, &user)
}

func CookieLoginHandler(ctx echo.Context) error {
	jwtCookie, err := ctx.Cookie(session.CookieName)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}

	var app types.App
	if err := db.Where("id = ?", ctx.(auth.Context).AppID).Find(&app).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, "server error signing")
	}
	signKey := app.SessionKey

	token, err := jwt.Parse(jwtCookie.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil
	})

	if err != nil || !token.Valid {
		return ctx.JSON(http.StatusUnauthorized, "invalid token:\n"+err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "claims error in token")
	}

	appID, okApp := claims["appId"].(string)
	userID, okUser := claims["userId"].(string)
	if !okApp || !okUser {
		return ctx.JSON(http.StatusUnauthorized, "Error converting UserID or AppID")
	}

	var user types.User
	user.ID = userID
	if err := db.Where("app_id = ?", appID).First(&user).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, &user)
}

func LogoutHandler(ctx echo.Context) error {
	// <-- Set user status to "CONNECTED" in the database. This is where memphis might be a good idea.

	ctx.SetCookie(&http.Cookie{
		Name:   session.CookieName,
		MaxAge: 0,
	})
	return ctx.NoContent(http.StatusOK)
}
