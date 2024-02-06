package handlers

import (
	"auth/database/types"
	"auth/service/middleware"
	"fmt"
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
func ListAppsHandler(ctx echo.Context) error {
	var apps []types.App
	if err := db.Find(&apps).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, &apps)
}

func GetAppHandler(ctx echo.Context) error {
	appID := ctx.(middleware.Context).AppID

	var app types.App
	app.ID = appID
	if err := db.Preload("Origins").Find(&app).Error; err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, "invalid app ID")
	}

	return ctx.JSON(http.StatusOK, &app)
}

func GetSecuritySettingsHandler(ctx echo.Context) error {
	appID := ctx.(middleware.Context).AppID

	dto := struct {
		Origins          []types.Origin `json:"allowedOrigins"`
		LockoutThreshold int            `json:"lockoutThreshold"`
		LockoutDuration  int            `json:"lockoutDuration"`
	}{}

	var app types.App
	app.ID = appID

	if err := db.Preload("Origins").First(&app).Error; err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, "invalid request")
	}

	copyFields(&app, &dto)

	return ctx.JSON(http.StatusOK, &dto)
}

func CreateAppHandler(ctx echo.Context) error {
	appCreateDTO := struct {
		Name string `json:"name"`
	}{}

	if err := ctx.Bind(&appCreateDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	app := types.App{Name: appCreateDTO.Name}
	if err := db.Create(&app).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusCreated, &app)
}

func DeleteAppHandler(ctx echo.Context) error {
	appID := ctx.(middleware.Context).AppID
	err := db.Delete(&types.App{}, appID).Error
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.NoContent(http.StatusNoContent)
}
