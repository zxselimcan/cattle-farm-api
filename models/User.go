package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UUID          string `gorm:"primaryKey;unique"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	FirstName     string
	LastName      string
	Email         string
	Password      string
	EmailVerified bool
	IsAdmin       bool
}
