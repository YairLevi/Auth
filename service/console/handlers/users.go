package handlers

import (
	"auth/service/database/types"
	auth "auth/service/middleware"
	"crypto/sha256"
	"encoding/base64"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ListUsersHandler(ctx echo.Context) error {
	appID := ctx.(auth.Context).AppID
	var app types.App
	app.ID = appID

	if err := db.Preload("Users").Find(&app).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, &app.Users)
}

func CreateUserHandler(ctx echo.Context) error {
	appID := ctx.(auth.Context).AppID

	var userDTO types.User
	userDTO.AppID = appID

	if err := ctx.Bind(&userDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	hashedPasswordInBytes := sha256.Sum256([]byte(userDTO.PasswordHash))
	hashedPasswordEncodedString := base64.StdEncoding.EncodeToString(hashedPasswordInBytes[:])
	userDTO.PasswordHash = hashedPasswordEncodedString

	if err := db.Create(&userDTO).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.NoContent(http.StatusCreated)
}

func DeleteUserHandler(ctx echo.Context) error {
	userID := ctx.Param("userId")
	user := types.User{}
	user.ID = userID
	if err := db.Delete(&user).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.NoContent(http.StatusNoContent)
}
