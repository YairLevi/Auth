package session

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	CookieName = "session"
)

type Payload struct {
	UserID string `json:"userId"`
}

type Config struct {
	Payload    interface{}
	Expiration time.Duration
	SigningKey string
}

func GenerateJWT(config Config) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iat"] = time.Now().Unix()
	claims["userId"] = config.Payload
	claims["exp"] = time.Now().Add(config.Expiration).Unix()
	secret := []byte(config.SigningKey)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
