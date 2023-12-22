package models

import (
	"time"
)

type User struct {
	ID              int
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Email           string    `json:"email"`
	Password        string    `json:"-"`
	EmailVerifiedAt time.Time `json:"-"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
	DeletedAt       time.Time `json:"-"`

	//gorm.Model
}
