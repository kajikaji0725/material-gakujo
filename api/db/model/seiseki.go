package model

import (
	"time"

	"github.com/szpp-dev-team/gakujo-api/model"
)

type Seiseki struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	UserID    int
	Seiseki   model.SeisekiRow `gorm:"embedded;embeddedPrefix:seiseki_"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
