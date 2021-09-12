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

func (controller *Controller) FetchUserInfoByName(gakujoUsername string) (*model.User, bool, error) {
	var userInfo model.User
	if err := controller.db.Table("users").Where("gakujo_username = ?", gakujoUsername).First(&userInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return &userInfo, true, nil
}

func (controller *Controller) UpdateSession(session, gakujoUsername string, expires time.Time) error {
	tx := controller.db.Begin()

	err := tx.Table("users").Where("gakujo_username = ?", gakujoUsername).Updates(map[string]interface{}{"session": session, "session_expires": expires}).Error
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

func (controller *Controller) FetchSeisekis(userID int) ([]model.Seiseki, error) {
	var seisekis []model.Seiseki
	if err := controller.db.Table("seisekis").Where("user_id = ?", userID).Find(&seisekis).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return seisekis, nil
}

// create first seiseki row. if duplicated, update update_at column
func (controller *Controller) CreateFirstSeiseki(gakujoSeiseki *gakujomodel.SeisekiRow, userID int) error {
	var gormSeiseki model.Seiseki
	var err error
	newGormSeiseki := model.Seiseki{
		UserID:    userID,
		Seiseki:   *gakujoSeiseki,
		CreatedAt: time.Now(),
	}

	err = controller.db.Table("seisekis").
		Where("user_id = ? AND seiseki_subject_name = ?", userID, gakujoSeiseki.SubjectName).
		First(&gormSeiseki).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	tx := controller.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err != nil {
		err = tx.Table("seisekis").Create(&newGormSeiseki).Error
		if err != nil {
			return err
		}
	} else {
		err = tx.Table("seisekis").
			Where("user_id = ? AND seiseki_subject_name = ?", userID, gakujoSeiseki.SubjectName).
			Update("updated_at", time.Now()).Error
		if err != nil {
			return err
		}
	}

	return nil
}
