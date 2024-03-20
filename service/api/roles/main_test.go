package roles

import (
	"auth/service/database"
	"auth/service/database/types"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

var (
	testServer *echo.Echo
)

func TestMain(m *testing.M) {
	testDb, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	database.PrepareDB(testDb)
	testDb.Create(&types.SecurityConfig{})

	db = testDb
	db.Create(&types.User{
		Email: "test@mail.com",
	})

	testServer = echo.New()

	m.Run()

	d, _ := testDb.DB()
	d.Close()
}
