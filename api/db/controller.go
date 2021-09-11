package db

import (
	"fmt"
	"time"

	"github.com/earlgray283/material-gakujo/api/db/model"
	"github.com/pkg/errors"
	gakujomodel "github.com/szpp-dev-team/gakujo-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Username string
	Password string
	DBname   string
	Port     string
}

type Controller struct {
	db *gorm.DB
}

func NewController(config *DBConfig) (*Controller, error) {
	db, err := gorm.Open(postgres.Open(dsn(config)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Seiseki{},
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Controller{db}, nil
}

func dsn(config *DBConfig) string {
	return fmt.Sprintf(
		"user=%s password=%s port=%s database=%s host=%s sslmode=disable",
		config.Username,
		config.Password,
		config.Port,
		config.DBname,
		config.Host,
	)
}

func (controller *Controller) FetchUserInfoBySessionID(sessionID string) (*model.User, error) {
	var userInfo model.User
	if err := controller.db.Table("users").Where("session = ?", sessionID).Find(&userInfo).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &userInfo, nil
}

func (controller *Controller) FetchUserInfoByName(gakujoUsername string) (*model.User, error) {
	var userInfo model.User
	if err := controller.db.Table("users").Where("gakujo_username = ?", gakujoUsername).Find(&userInfo).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &userInfo, nil
}

func (controller *Controller) UpdateSession(session, gakujoUsername string, expires time.Time) error {
	tx := controller.db.Begin()
	err := tx.Table("users").Where("gakujo_username = ?", gakujoUsername).Update("session", session).Update("session_expires", expires).Error
	if err != nil {
		tx.Rollback()
		return errors.WithStack(err)
	}
	tx.Commit()
	return nil
}

func (controller *Controller) CreateUser(user *model.User) error {
	tx := controller.db.Begin()
	err := tx.Table("users").Create(user).Error
	if err != nil {
		tx.Rollback()
		return errors.WithStack(err)
	}
	tx.Commit()
	return nil
}

func (controller *Controller) FetchSeisekis(gakujoUsername string) ([]gakujomodel.SeisekiRow, error) {
	var seisekis []model.Seiseki
	if err := controller.db.Table("users").Where("gakujo_username = ?", gakujoUsername).Find(&seisekis).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	var gakujoSeisekis []gakujomodel.SeisekiRow
	for _, seiseki := range seisekis {
		gakujoSeisekis = append(gakujoSeisekis, seiseki.Seiseki)
	}

	return gakujoSeisekis, nil

}

func (controller *Controller) CreateSeisekis(gakujoSeisekis []*gakujomodel.SeisekiRow, userID int) error {
	var gormSeisekis []model.Seiseki
	for _, gakujoSeiseki := range gakujoSeisekis {
		gormSeiseki := model.Seiseki{
			UserID:    userID,
			Seiseki:   *gakujoSeiseki,
			CreatedAt: time.Now(),
		}
		gormSeisekis = append(gormSeisekis, gormSeiseki)
	}

	tx := controller.db.Begin()
	err := tx.Table("seisekis").Create(gormSeisekis).Error
	if err != nil {
		tx.Rollback()
		return errors.WithStack(err)
	}
	tx.Commit()
	return nil
}
