package types

import (
	"time"
)

type User struct {
	Model
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"passwordHash"`
	PhoneNumber  string    `json:"phoneNumber"`
	PhotoURL     string    `json:"photoURL"`
	LastLogin    time.Time `json:"lastLogin"`
	Birthday     time.Time `json:"birthday"`

	// this is redundant, but is used to enable cascade delete
	UserRoles []UserRole `gorm:"constraint:OnDelete:CASCADE"`
}

type EmailFilter struct {
	Model
	Email       string `json:"email"`
	IsWhitelist bool   `json:"isWhitelist"`

	SecurityConfig   SecurityConfig
	SecurityConfigID string
}

type SecurityConfig struct {
	Model
	Origins          []Origin      `json:"origins"`
	LockoutThreshold int           `json:"lockoutThreshold" gorm:"default:5"`
	LockoutDuration  int           `json:"lockoutDuration" gorm:"default:30"`
	SessionKey       string        `json:"sessionKey" gorm:"default:'session key'"`
	EmailFilters     []EmailFilter `json:"emailFilters"`
}

type Origin struct {
	Model
	URL string `json:"url"`

	SecurityConfig   SecurityConfig
	SecurityConfigID string
}

type Role struct {
	Model
	Name string `json:"name" gorm:"unique"`
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
