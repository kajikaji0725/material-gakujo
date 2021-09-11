package server

import (
	"log"
	"net/http"

	"github.com/earlgray283/material-gakujo/api/server/crawle"
)

func (api *ApiServer) Crawle(rw http.ResponseWriter, r *http.Request) {
	session, _ := r.Cookie("GAKUJO_SESSION")
	user, err := api.controller.FetchUserInfoBySessionID(session.Value)
	if err != nil {
		log.Println(err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	username, password, err := crawle.AuthInfoFromUser(user, api.cryptoKey)
	if err != nil {
		log.Println(err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	crawler, err := crawle.NewCrawler(api.controller, username, password, user.ID)
	if err != nil {
		log.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := crawler.CrawleAll(); err != nil {
		log.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
