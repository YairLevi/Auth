package handlers

import (
	"auth/service/database/types"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetLockoutThresholdHandler(ctx echo.Context) error {
	dto := struct {
		Threshold int `json:"threshold"`
	}{}
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	var security types.SecurityConfig
	db.First(&security)
	security.LockoutThreshold = dto.Threshold
	if err := db.Save(&security).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error saving new lockout threshold")
	}
	return ctx.NoContent(http.StatusOK)
}

func SetLockoutDurationHandler(ctx echo.Context) error {
	dto := struct {
		Duration int `json:"duration"`
	}{}
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	var security types.SecurityConfig
	db.First(&security)
	security.LockoutDuration = dto.Duration
	if err := db.Save(&security).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error saving new lockout duration")
	}
	return ctx.NoContent(http.StatusOK)
}

func SetSessionKeyHandler(ctx echo.Context) error {
	dto := struct {
		SessionKey string `json:"sessionKey"`
	}{}
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	var security types.SecurityConfig
	db.First(&security)
	security.SessionKey = dto.SessionKey
	if err := db.Save(&security).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error saving new session key")
	}
	return ctx.NoContent(http.StatusOK)
}

func AddOriginHandler(ctx echo.Context) error {
	dto := struct {
		Origin string `json:"origin"`
	}{}
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	origin := types.Origin{
		URL: dto.Origin,
	}
	if err := db.Create(&origin).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.NoContent(http.StatusCreated)
}
