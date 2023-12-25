package handlers

import (
	"auth-service/database/types"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ListAppsHandler(ctx echo.Context) error {
	var apps []types.App
	if err := db.Find(&apps).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, &apps)
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
	appID := ctx.Param("appId")
	err := db.Delete(&types.App{}, appID).Error
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.NoContent(http.StatusNoContent)
}
