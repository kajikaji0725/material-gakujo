package auth

import (
	"encoding/hex"
	"net/http"
	"time"

	"github.com/earlgray283/material-gakujo/api/db"
	"github.com/earlgray283/material-gakujo/api/db/model"
	"github.com/gorilla/securecookie"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type SessionController struct {
	controller *db.Controller
}

func NewSessionController(controller *db.Controller) *SessionController {
	return &SessionController{
		controller: controller,
	}
}

func (sc *SessionController) GenerateNewSession(gakujoUsername string) (*http.Cookie, error) {
	session := generateSessionID()
	expires := time.Now().Add(7 * 24 * time.Hour)
	if err := sc.controller.UpdateSession(session, gakujoUsername, expires); err != nil {
		return nil, err
	}

	cookie := http.Cookie{
		Name:    "GAKUJO_SESSION",
		Value:   session,
		Expires: expires,
		//Secure:   true,
		HttpOnly: true,
	}

	return &cookie, nil
}

func (sc *SessionController) RemoveSession(gakujoUsername string) (*http.Cookie, error) {
	if err := sc.controller.UpdateSession("", gakujoUsername, time.Now()); err != nil {
		return nil, err
	}

	cookie := http.Cookie{
		Name: "GAKUJO_SESSION",
		//Secure:   true,
		HttpOnly: true,
		Path:     "/",
	}

	return &cookie, nil
}

func (sc *SessionController) CheckSession(sessionID string) (*model.User, bool, error) {
	user, err := sc.controller.FetchUserInfoBySessionID(sessionID)
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

func generateSessionID() string {
	b := securecookie.GenerateRandomKey(11)
	return hex.EncodeToString(b)
}

func NewRemovedCookie() *http.Cookie {
	cookie := http.Cookie{
		Name:    "GAKUJO_SESSION",
		Expires: time.Now(),
		//Secure:   true,
		HttpOnly: true,
	}

	return &cookie
}
