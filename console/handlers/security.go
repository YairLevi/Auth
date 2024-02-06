package handlers

import (
	"auth/database/types"
	auth "auth/service/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetLockoutThresholdHandler(ctx echo.Context) error {
	appID := ctx.(auth.Context).AppID
	dto := struct {
		Threshold int `json:"threshold"`
	}{}

	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	var app types.App
	app.ID = appID
	db.First(&app)
	app.LockoutThreshold = dto.Threshold
	if err := db.Save(&app).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error saving new lockout threshold")
	}
	return ctx.NoContent(http.StatusOK)
}

func SetLockoutDurationHandler(ctx echo.Context) error {
	appID := ctx.(auth.Context).AppID
	dto := struct {
		Duration int `json:"duration"`
	}{}

	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	var app types.App
	app.ID = appID
	db.First(&app)

	app.LockoutDuration = dto.Duration
	if err := db.Save(&app).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error saving new lockout duration")
	}
	return ctx.NoContent(http.StatusOK)
}

func SetSessionKeyHandler(ctx echo.Context) error {
	appID := ctx.(auth.Context).AppID
	dto := struct {
		Key string `json:"key"`
	}{}

	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	var app types.App
	app.ID = appID
	db.First(&app)

	app.SessionKey = dto.Key
	if err := db.Save(&app).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error saving new session key")
	}
	return ctx.NoContent(http.StatusOK)
}

func AddOriginHandler(ctx echo.Context) error {
	appID := ctx.(auth.Context).AppID
	dto := struct {
		Origin string `json:"origin"`
	}{}

	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := db.Create(&types.Origin{URL: dto.Origin, AppID: appID}).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.NoContent(http.StatusCreated)
}
