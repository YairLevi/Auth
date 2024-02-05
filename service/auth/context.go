package auth

import "github.com/labstack/echo/v4"

type Context struct {
	echo.Context
	AppID string `json:"appId"`
}
