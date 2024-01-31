package types

import (
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
	Name             string   `json:"name"`
	Users            []User   `json:"users"`
	Origins          []Origin `json:"allowedOrigins"`
	LockoutThreshold int      `json:"lockoutThreshold" gorm:"default:5"`
	LockoutDuration  int      `json:"lockoutDuration" gorm:"default:30"`
	SessionKey       string   `json:"sessionKey" gorm:"default:'session key'"`
}

type Origin struct {
	Model
	AppID string `json:"-"`
	URL   string `json:"url"`
}
