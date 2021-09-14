package model

import (
	"time"

	"github.com/szpp-dev-team/gakujo-api/model"
)

type ClassNotice struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	UserID      int
	ClassNotice model.SeisekiRow `gorm:"embedded;embeddedPrefix:class_notice_"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
