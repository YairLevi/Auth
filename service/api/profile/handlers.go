package profile

import (
	"auth/service/database"
	"auth/service/database/types"
	auth "auth/service/middleware"
	"auth/service/tools/password"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

var db = database.DB

type PasswordReset struct {
}

func UpdateHandler(ctx echo.Context) error {
	var user types.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, "error in payload password reset")
	}
	user.AppID = ctx.(auth.Context).AppID
	return nil
}

func ResetPasswordHandler(ctx echo.Context) error {
	dto := struct {
		UserID      string `json:"userId"`
		NewPassword string `json:"newPassword"`
	}{}
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, "error in payload password reset")
	}

	appID := ctx.(auth.Context).AppID
	userID := dto.UserID

	var user types.User
	user.ID = dto.UserID
	user.AppID = appID
	if err := db.Where(&user).Find(&user).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, fmt.Sprint("invalid user ID or app ID:", userID, appID))
	}

	user.PasswordHash = password.Encrypt(dto.NewPassword)
	if err := db.Save(&user).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, "failed to save user, password reset")
	}

	return ctx.NoContent(http.StatusCreated)
}
