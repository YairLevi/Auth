package session

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	CookieName = "session"
)

type Payload struct {
	AppID  string `json:"appId"`
	UserID string `json:"userId"`
}

type Config struct {
	Payload    Payload
	Expiration time.Duration
	SigningKey string
}

func GenerateJWT(config Config) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iat"] = time.Now().Unix()
	claims["appId"] = config.Payload.AppID
	claims["userId"] = config.Payload.UserID
	claims["exp"] = time.Now().Add(config.Expiration).Unix()
	secret := []byte(config.SigningKey)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
