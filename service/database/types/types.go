package types

import (
	"time"
)

type User struct {
	Model
	App          App
	AppID        string    `json:"-"`
	Username     string    `json:"username"`
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
	App   App
	AppID string `json:"-"`
	URL   string `json:"url"`
}

type Role struct {
	Model
	App   App    `json:"-"`
	AppID string `json:"appId"`
	Name  string `json:"name" gorm:"unique"`

	// this is redundant, but is used to enable a cascade delete from children to parents.
	UserRoles []UserRole `gorm:"constraint:OnDelete:CASCADE"`
}

type UserRole struct {
	Model
	RoleID string
	Role   Role
	UserID string
	User   User
}
