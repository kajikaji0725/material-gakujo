package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/earlgray283/material-gakujo/api/db"
	auth "github.com/earlgray283/material-gakujo/api/server/libauth"
)

func AuthMiddleware(controller *db.Controller) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			session, err := r.Cookie("GAKUJO_SESSION")
			if err != nil {
				log.Println(r.URL.Path)
				if r.URL.Path == "/api/auth/login" || r.URL.Path == "/api/auth/register" {
					h.ServeHTTP(rw, r)
					return
				}
				http.Error(rw, "Please login", http.StatusUnauthorized)
				return
			}

			user, ok, err := auth.CheckSession(controller, session.Value)
			if err != nil {
				http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			if !ok {
				http.Error(rw, "Session is not valid. Please login again.", http.StatusUnauthorized)
				return
			}

			// セッションが残っているなら login はパスすることができる
			if r.URL.Path == "/api/auth/login" {
				rw.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(rw).Encode(user)
				return
			}

			r.Header.Set("gakujo_username", user.Username)

			h.ServeHTTP(rw, r)
		})

	}
}

func LoggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s\n", r.Method, r.RequestURI)
		log.Printf("Content-Type: %s\n", r.Header.Get("Content-Type"))
		if r.Method == http.MethodPost || r.Method == http.MethodPut {
			log.Println("Payload:")
			if r == nil {
				fmt.Println("WTF?!")
				return
			}
			b, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println(err)
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
			r.Body = io.NopCloser(bytes.NewReader(b))
			log.Println(string(b))
		}

		h.ServeHTTP(rw, r)
	})
}
