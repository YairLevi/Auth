package handlers

import (
	"auth-service/database/types"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func AddOriginHandler(ctx echo.Context) error {
	appID := ctx.Param("appId")
	if appID == "" {
		return ctx.JSON(http.StatusBadRequest, "invalid app ID")
	}
	dto := struct {
		Origin string `json:"origin"`
	}{}

	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := db.Create(&types.AllowedOrigins{URL: dto.Origin, AppID: appID}).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.NoContent(http.StatusCreated)
}

func GetOriginsHandler(ctx echo.Context) error {
	appID := ctx.Param("appId")
	if appID == "" {
		return ctx.JSON(http.StatusBadRequest, "invalid app ID")
	}

	var origins []types.AllowedOrigins
	err := db.Where("app_id = ?", appID).Find(&origins).Error
	fmt.Println(origins)
	if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		return ctx.JSON(http.StatusInternalServerError, "")
	}

	return ctx.JSON(http.StatusOK, &origins)
}
