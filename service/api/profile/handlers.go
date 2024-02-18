package profile

import (
	"auth/service/database"
	"auth/service/database/types"
	"github.com/labstack/echo/v4"
	"net/http"
)

var db = database.DB

func Update(ctx echo.Context) error {
	dto := struct {
		PhotoURL string `json:"photoURL"`
		Username string `json:"username"`
		UserID   string `json:"userId"`
	}{}
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, "error in payload password reset")
	}

	var user types.User
	user.ID = dto.UserID
	if err := db.Where(&user).Find(&user).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, "error in fetching user")
	}

	if dto.PhotoURL != "" {
		user.PhotoURL = dto.PhotoURL
	}
	if dto.Username != "" {
		user.Username = dto.Username
	}

	if err := db.Save(&user).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, "error saving user")
	}
	return ctx.JSON(http.StatusCreated, &user)
}

//func ResetPasswordHandler(ctx echo.Context) error {
//	dto := struct {
//		UserID      string `json:"userId"`
//		NewPassword string `json:"newPassword"`
//	}{}
//	if err := ctx.Bind(&dto); err != nil {
//		return ctx.JSON(http.StatusBadRequest, "error in payload password reset")
//	}
//
//	appID := ctx.(auth.Context).AppID
//	userID := dto.UserID
//
//	var user types.User
//	user.ID = dto.UserID
//	user.AppID = appID
//	if err := db.Where(&user).Find(&user).Error; err != nil {
//		return ctx.JSON(http.StatusBadRequest, fmt.Sprint("invalid user ID or app ID:", userID, appID))
//	}
//
//	user.PasswordHash = password.Encrypt(dto.NewPassword)
//	if err := db.Save(&user).Error; err != nil {
//		return ctx.JSON(http.StatusBadRequest, "failed to save user, password reset")
//	}
//
//	return ctx.NoContent(http.StatusCreated)
//}
