package server

import (
	"crypto/subtle"
	"log"
	"net/http"
	"time"

	"github.com/earlgray283/material-gakujo/api/db/model"
	auth "github.com/earlgray283/material-gakujo/api/server/libauth"
	"github.com/szpp-dev-team/gakujo-api/gakujo"
)

func (api *ApiServer) Login(rw http.ResponseWriter, r *http.Request) {
	username := r.FormValue("gakujo_username")
	password := r.FormValue("gakujo_password")

	if username == "" || password == "" {
		http.Error(rw, "username or password is empty", http.StatusBadRequest)
		return
	}

	user, ok, err := api.controller.FetchUserInfoByName(username)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(rw, "No such username.", http.StatusUnauthorized)
		return
	}

	gakujoDecryptedPassword, err := auth.Decrypt(user.GakujoEncryptedPassword, api.cryptoKey)
	if err != nil {
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if subtle.ConstantTimeCompare([]byte(password), gakujoDecryptedPassword) != 1 {
		log.Println(password, gakujoDecryptedPassword)
		http.Error(rw, "username or password is invalid", http.StatusUnauthorized)
		return
	}

	sessionValue := auth.GenSessionID()
	expires := time.Now().Add(7 * 24 * time.Hour)
	if err := api.controller.UpdateSession(string(sessionValue), username, expires); err != nil {
		log.Println(err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	sessionCookie := auth.NewCookie(string(sessionValue), expires)
	http.SetCookie(rw, sessionCookie)
}

func (api *ApiServer) Logout(rw http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("GAKUJO_SESSION")
	user, _ := api.controller.FetchUserInfoBySessionID(cookie.Value)

	if err := api.controller.UpdateSession("", user.GakujoUsername, time.Now()); err != nil {
		log.Println(err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, _ = rw.Write([]byte("logout"))
}

func (api *ApiServer) RegistNewUser(rw http.ResponseWriter, r *http.Request) {
	gakujoUsername := r.FormValue("gakujo_username")
	gakujoPassword := r.FormValue("gakujo_password")
	username := r.FormValue("username")
	email := r.FormValue("email")

	if gakujoPassword == "" || gakujoUsername == "" || username == "" || email == "" {
		http.Error(rw, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := checkGakujoUser(gakujoUsername, gakujoPassword); err != nil {
		log.Println(err)
		http.Error(rw, "Could not login to gakujo. You might have set wrong user info.", http.StatusUnauthorized)
		return
	}

	encryptedGakujoPassword, err := auth.Encrypt([]byte(gakujoPassword), api.cryptoKey)
	if err != nil {
		log.Println(err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	user := &model.User{
		GakujoUsername:          gakujoUsername,
		GakujoEncryptedPassword: encryptedGakujoPassword,
		Username:                username,
		Email:                   email,
		CreatedAt:               time.Now(),
	}
	if err := api.controller.CreateUser(user); err != nil {
		log.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	sessionValue := auth.GenSessionID()
	expires := time.Now().Add(7 * 24 * time.Hour)
	if err := api.controller.UpdateSession(string(sessionValue), username, expires); err != nil {
		log.Println(err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	sessionCookie := auth.NewCookie(string(sessionValue), expires)
	http.SetCookie(rw, sessionCookie)

	_, _ = rw.Write([]byte("success"))
}

func checkGakujoUser(username, password string) error {
	c := gakujo.NewClient()
	return c.Login(username, password)
}
