package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (api *ApiServer) FetchSeiseki(rw http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.Header.Get("user_id"))
	seisekis, err := api.controller.FetchSeisekis(userID)
	if err != nil {
		log.Println(err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(rw).Encode(seisekis)
}
