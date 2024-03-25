package standard

import (
	"auth/service/database/types"
	"auth/service/session"
	"auth/service/tools/password"
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

	user := registerDTO.User
	user.PasswordHash = password.Encrypt(registerDTO.Password)

	if err := db.Create(&user).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, &user)
}

func EmailPasswordLoginHandler(ctx echo.Context) error {
	loginDTO := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	if err := ctx.Bind(&loginDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if IsLocked(loginDTO.Email) {
		return ctx.JSON(http.StatusForbidden, fmt.Sprintf("in lockout state for a few seconds."))
	}
	var user types.User
	if err := db.Where("email = ?", loginDTO.Email).First(&user).Error; err != nil {
		return ctx.JSON(http.StatusUnauthorized, "unauthorized")
	}

	if !password.IsEqual(loginDTO.Password, user.PasswordHash) {
		Fail(loginDTO.Email)
		return ctx.JSON(http.StatusUnauthorized, "unauthorized")
	}
	Succeed(loginDTO.Email)

	var security types.SecurityConfig
	if err := db.First(&security).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	jwtCookie, err := session.GenerateJWT(session.Config{
		Expiration: time.Hour * 24,
		SigningKey: security.SessionKey,
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
		Expires:  time.Now().Add(time.Hour * 24),
		Path:     "/",
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

	var security types.SecurityConfig
	if err := db.First(&security).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, "server error signing")
	}

	token, err := jwt.Parse(jwtCookie.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(security.SessionKey), nil
	})

	if err != nil || !token.Valid {
		return ctx.JSON(http.StatusUnauthorized, "invalid token:\n"+err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "claims error in token")
	}

	userID, okUser := claims["userId"].(string)
	if !okUser {
		return ctx.JSON(http.StatusUnauthorized, "Error converting UserID")
	}

	var user types.User
	user.ID = userID
	if err := db.Where(&user).First(&user).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, &user)
}

func LogoutHandler(ctx echo.Context) error {
	// <-- Set user status to "CONNECTED" in the database. This is where memphis might be a good idea.

	ctx.SetCookie(&http.Cookie{
		Name:   session.CookieName,
		Path:   "/",
		MaxAge: 0,
	})
	return ctx.NoContent(http.StatusOK)
}
