package standard

import (
	"auth-service/database"
	"auth-service/database/types"
	mjwt "auth-service/service/jwt"
	"crypto/sha256"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

var db = database.DB

func RegisterHandler(ctx echo.Context) error {
	registerDTO := struct {
		types.User
		Password string `json:"password"`
		AppID    string `json:"appId"`
	}{}

	if err := ctx.Bind(&registerDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	hashedPasswordInBytes := sha256.Sum256([]byte(registerDTO.Password))
	hashedPasswordEncodedString := base64.StdEncoding.EncodeToString(hashedPasswordInBytes[:])
	user := registerDTO.User
	user.PasswordHash = hashedPasswordEncodedString
	user.AppID = registerDTO.AppID

	if err := db.Create(&user).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)

	}

	return ctx.JSON(http.StatusOK, &user)
}

func EmailPasswordLoginHandler(ctx echo.Context) error {
	loginDTO := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		AppID    string `json:"appId"`
	}{}

	if err := ctx.Bind(&loginDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	var user types.User
	res := db.Where("email = ? AND app_id = ?", loginDTO.Email, loginDTO.AppID).First(&user)
	if res.Error != nil {
		return ctx.JSON(http.StatusUnauthorized, "unauthorized")
	}

	hashedPasswordInBytes := sha256.Sum256([]byte(loginDTO.Password))
	hashedPasswordEncodedString := base64.StdEncoding.EncodeToString(hashedPasswordInBytes[:])

	if hashedPasswordEncodedString != user.PasswordHash {
		return ctx.JSON(http.StatusUnauthorized, "unauthorized")
	}
	jwtCookie, err := mjwt.GenerateJWT(mjwt.Config{
		Expiration: time.Hour * 24,
		SigningKey: mjwt.SecretKey,
		Payload: mjwt.Payload{
			AppID:  loginDTO.AppID,
			UserID: user.ID,
		},
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.SetCookie(&http.Cookie{
		Name:     mjwt.CookieName,
		Value:    jwtCookie,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		Secure:   true,
	})

	return ctx.JSON(http.StatusOK, &user)
}

func CookieLoginHandler(ctx echo.Context) error {
	jwtCookie, err := ctx.Cookie(mjwt.CookieName)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}

	token, err := jwt.Parse(jwtCookie.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(mjwt.SecretKey), nil
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
		return ctx.JSON(http.StatusBadRequest, "Error converting UserID or AppID")
	}

	var user types.User
	user.ID = userID
	if err := db.Where("app_id = ?", appID).First(&user).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, &user)
}

func LogoutHandler(ctx echo.Context) error {
	logoutDTO := struct {
		AppID string `json:"appId"`
	}{}
	if err := ctx.Bind(&logoutDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	// <-- Set user status to "CONNECTED" in the database. This is where memphis might be a good idea.

	ctx.SetCookie(&http.Cookie{
		Path:   "/",
		Name:   mjwt.CookieName,
		MaxAge: 0,
	})
	return ctx.NoContent(http.StatusOK)
}
