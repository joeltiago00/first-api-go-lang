package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Password        string
	EmailVerifiedAt time.Time
	gorm.Model
}
