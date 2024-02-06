package middleware

import (
	"auth/database"
	"auth/database/types"
	"github.com/labstack/echo/v4"
	"net/http"
)

const ApplicationHeader = "X-App-ID"

type Context struct {
	echo.Context
	AppID string `json:"appId"`
}

func AppID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		appID := ctx.Request().Header.Get(ApplicationHeader)
		var app types.App
		if err := database.DB.Where("id = ?", appID).Find(&app).Error; err != nil {
			return ctx.JSON(http.StatusUnauthorized, "invalid app id")
		}

		return next(Context{
			Context: ctx,
			AppID:   appID,
		})
	}
}
