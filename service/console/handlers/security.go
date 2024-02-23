package handlers

import (
	"auth/service/database/types"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"reflect"
)

func copyFields(originalPtr interface{}, partialPtr interface{}) {
	originalValue := reflect.ValueOf(originalPtr).Elem()
	partialValue := reflect.ValueOf(partialPtr).Elem()

	for i := 0; i < originalValue.NumField(); i++ {
		fieldName := originalValue.Type().Field(i).Name
		if partialValue.FieldByName(fieldName).IsValid() {
			partialValue.FieldByName(fieldName).Set(originalValue.Field(i))
		}
	}
}

func GetSecuritySettingsHandler(ctx echo.Context) error {
	dto := struct {
		Origins          []types.Origin `json:"allowedOrigins"`
		LockoutThreshold int            `json:"lockoutThreshold"`
		LockoutDuration  int            `json:"lockoutDuration"`
	}{}

	var security types.SecurityConfig
	err := db.Preload("Origins").First(&security).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		db.Create(&security)
	} else if err != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid request")
	}
	// because when transferred, is null. golang stuff...
	if len(security.Origins) == 0 {
		security.Origins = make([]types.Origin, 0)
	}

	copyFields(&security, &dto)
	return ctx.JSON(http.StatusOK, &dto)
}

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
	var security types.SecurityConfig
	if err := db.First(&security).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	origin := types.Origin{
		URL:            dto.Origin,
		SecurityConfig: security,
	}
	if err := db.Create(&origin).Error; err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.NoContent(http.StatusCreated)
}

func RemoveOriginHandler(ctx echo.Context) error {
	originID := ctx.Param("originId")
	if err := db.Where("id = ?", originID).Delete(&types.Origin{}).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.NoContent(http.StatusOK)
}
