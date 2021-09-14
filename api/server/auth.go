package server

import (
	"crypto/subtle"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/earlgray283/material-gakujo/api/db/model"
	auth "github.com/earlgray283/material-gakujo/api/server/libauth"
	"github.com/szpp-dev-team/gakujo-api/gakujo"
)

func (api *ApiServer) Login(rw http.ResponseWriter, r *http.Request) {
	gakujoUsername := r.FormValue("gakujo_username")
	gakujoPassword := r.FormValue("gakujo_password")

	if gakujoUsername == "" || gakujoPassword == "" {
		http.Error(rw, "username or password is empty", http.StatusBadRequest)
		return
	}

	user, ok, err := api.controller.FetchUserInfoByName(gakujoUsername)
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
	if subtle.ConstantTimeCompare([]byte(gakujoPassword), gakujoDecryptedPassword) != 1 {
		log.Println(gakujoPassword, gakujoDecryptedPassword)
		http.Error(rw, "username or password is invalid", http.StatusUnauthorized)
		return
	}

	sessionCookie, err := api.sessionController.GenerateNewSession(gakujoUsername)
	if err != nil {
		log.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
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
		log.Println("invalid input")
		http.Error(rw, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := checkGakujoUser(gakujoUsername, gakujoPassword); err != nil {
		log.Println("gakujo login failed")
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

	/*
		sessionCookie, err := api.sessionController.GenerateNewSession(gakujoUsername)
		if err != nil {
			log.Println(err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		http.SetCookie(rw, sessionCookie)
	*/

	_, _ = rw.Write([]byte("success"))
}

// check session and return user info
func (api *ApiServer) FetchUser(rw http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("GAKUJO_SESSION")
	session := cookie.Value

	user, err := api.controller.FetchUserInfoBySessionID(session)
	if err != nil {
		log.Println(err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(rw).Encode(user)
}

func checkGakujoUser(username, password string) error {
	c := gakujo.NewClient()
	return c.Login(username, password)
}
