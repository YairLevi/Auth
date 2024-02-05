package middleware

import (
	"auth/database"
	"auth/database/types"
	"auth/service/auth"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AppID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		appID := ctx.Param("appId")
		var app types.App
		if err := database.DB.Where("id = ?", appID).Find(&app).Error; err != nil {
			return ctx.JSON(http.StatusUnauthorized, "invalid app id")
		}

		return next(auth.Context{
			Context: ctx,
			AppID:   appID,
		})
	}
}
