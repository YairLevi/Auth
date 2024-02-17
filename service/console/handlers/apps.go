package handlers

import (
	"auth/service/database/types"
	"github.com/labstack/echo/v4"
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
	if err := db.Preload("Origins").First(&security).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid request")
	}

	copyFields(&security, &dto)
	return ctx.JSON(http.StatusOK, &dto)
}
