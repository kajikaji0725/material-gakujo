package auth

import (
	"encoding/hex"
	"time"

	"github.com/earlgray283/material-gakujo/api/db"
	"github.com/earlgray283/material-gakujo/api/db/model"
	"github.com/gorilla/securecookie"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func CheckSession(controller *db.Controller, sessionID string) (*model.User, bool, error) {
	user, err := controller.FetchUserInfoBySessionID(sessionID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, err
	}

	// セッション有効期限切れ
	if time.Now().After(user.SessionExpires) {
		return nil, false, nil
	}

	return user, true, nil
}

func GenSessionID() string {
	b := securecookie.GenerateRandomKey(11)
	return hex.EncodeToString(b)
}
