package handlers

import (
	"auth/service/database/types"
	"auth/service/tools/password"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func ListUsersHandler(ctx echo.Context) error {
	users := make([]types.User, 0)
	if err := db.Find(&users).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, &users)
}

func CreateUserHandler(ctx echo.Context) error {
	var userDTO types.User
	if err := ctx.Bind(&userDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	// TODO: password hash should not be in the user DTO object, but passed in different format ?...
	userDTO.PasswordHash = password.Encrypt(userDTO.PasswordHash)

	if err := db.Create(&userDTO).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.NoContent(http.StatusCreated)
}

func DeleteUserHandler(ctx echo.Context) error {
	user := types.User{}
	user.ID = ctx.Param("userId")
	if err := db.Delete(&user).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.NoContent(http.StatusNoContent)
}

func UpdateUserHandler(ctx echo.Context) error {
	var user types.User
	user.ID = ctx.Param("userId")
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	var temp types.User
	if err := db.Where("id = ?", user.ID).Find(&temp).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.JSON(http.StatusBadRequest, "error, no such user.")
	}

	if err := db.Save(&user).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, fmt.Errorf("error saving user %v", err))
	}

	return ctx.NoContent(http.StatusNoContent)
}
