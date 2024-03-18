package standard

import (
	"auth/service/database"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"testing"
)

var (
	testServer *echo.Echo
)

func TestMain(m *testing.M) {
	log.Println("Before testing")
	testDb, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database.PrepareDB(testDb)

	db = testDb
	testServer = echo.New()

	m.Run()
	log.Println("Post testing")
	d, _ := testDb.DB()
	d.Close()
}
