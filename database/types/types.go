package types

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Model
	AppID        string    `json:"-"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"passwordHash"`
	PhoneNumber  string    `json:"phoneNumber"`
	LastLogin    time.Time `json:"lastLogin"`
	Birthday     time.Time `json:"birthday"`
}

type App struct {
	Model
	Name    string           `json:"name"`
	Users   []User           `json:"users"`
	Origins []AllowedOrigins `json:"allowedOrigins"`
}

type AllowedOrigins struct {
	gorm.Model
	URL   string `json:"url"`
	AppID string `json:"appId"`
}
