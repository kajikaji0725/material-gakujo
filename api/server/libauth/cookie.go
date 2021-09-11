package auth

import (
	"net/http"
	"time"
)

func NewCookie(session string, expires time.Time) *http.Cookie {
	sessionCookie := http.Cookie{
		Name:     "GAKUJO_SESSION",
		Value:    session,
		Expires:  expires,
		Secure:   true,
		HttpOnly: true,
	}

	return &sessionCookie
}
