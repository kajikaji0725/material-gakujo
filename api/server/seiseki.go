package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (api *ApiServer) FetchSeiseki(rw http.ResponseWriter, r *http.Request) {
	gakujoUsername := r.Header.Get("gakujo_username")
	seisekis, err := api.controller.FetchSeisekis(gakujoUsername)
	if err != nil {
		log.Println(err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(rw).Encode(seisekis)
}
