package model

import (
	"time"
)

type User struct {
	ID                      int `gorm:"primaryKey;autoIncrement"`
	GakujoUsername          string
	GakujoEncryptedPassword string `json:"-"`
	Username                string
	Email                   string
	Grade                   int
	Session                 string
	SessionExpires          time.Time
	CreatedAt               time.Time
	UpdatedAt               time.Time
}
