package model

import (
	"time"
)

type User struct {
	ID                      int       `gorm:"primaryKey;autoIncrement" json:"id"`
	GakujoUsername          string    `json:"gakujo_username"`
	GakujoEncryptedPassword string    `json:"-"`
	Username                string    `json:"username"`
	Email                   string    `json:"email"`
	Grade                   int       `json:"grade"`
	Session                 string    `json:"-"`
	SessionExpires          time.Time `json:"-"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}
