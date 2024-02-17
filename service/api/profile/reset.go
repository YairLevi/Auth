package profile

//
//import (
//	"auth/service/database/types"
//	"auth/service/session"
//	"fmt"
//	"github.com/labstack/echo/v4"
//	"github.com/labstack/gommon/log"
//	"net/http"
//	"net/smtp"
//	"time"
//)
//
//type ResetPayload struct {
//	Email string `json:"email"`
//}
//
//func RequestPasswordReset(ctx echo.Context) error {
//	dto := struct {
//		Email string `json:"email"`
//	}{}
//	if err := ctx.Bind(&dto); err != nil {
//		return ctx.JSON(http.StatusBadRequest, "error in payload")
//	}
//	var user types.User
//	user.ID = dto.Email
//	if err := db.Where(&user).Find(&user).Error; err != nil {
//		return ctx.JSON(http.StatusBadRequest, "invalid user")
//	}
//	var security types.SecurityConfig
//	if err := db.First(&security).Error; err != nil {
//		return ctx.JSON(http.StatusInternalServerError, "error in request password")
//	}
//
//	token, err := session.GenerateJWT(session.Config{
//		Expiration: time.Minute,
//		SigningKey: security.SessionKey,
//		Payload: ResetPayload{
//			Email: dto.Email,
//		},
//	})
//	if err != nil {
//		return ctx.JSON(http.StatusInternalServerError, "error creating token for reset")
//	}
//
//	sendResetEmail(dto.Email, token)
//	return ctx.NoContent(http.StatusNoContent)
//}
//
//func ChangePassword(ctx echo.Context) error {
//
//}
//
//func sendResetEmail(to string, token string) {
//	from := "someemail@support.com"
//	password := "the password"
//
//	smtpHost := "smtp.gmail.com"
//	smtpPort := "587"
//
//	message := fmt.Sprint("")
//
//	message := []byte("Reset password at: <a href=''>")
//	auth := smtp.PlainAuth("", from, password, smtpHost)
//	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
//	if err != nil {
//		log.Error(err)
//		return
//	}
//}
