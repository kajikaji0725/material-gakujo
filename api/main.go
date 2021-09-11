package main

import (
	"log"
	"net/http"
	"os"

	"github.com/earlgray283/material-gakujo/api/db"
	"github.com/earlgray283/material-gakujo/api/server"
)

func main() {
	config := db.DBConfig{
		Host:     "material-gakujo-db",
		Username: "root",
		Password: "root",
		DBname:   "root",
		Port:     "5432",
	}
	cryptoKey := os.Getenv("CRYPTO_KEY")
	if len(cryptoKey) != 32 {
		log.Fatal("The length of CRYPTO_KEY must be 32")
	}
	apiServer, err := server.NewApiServer(&config, []byte(cryptoKey))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server Started")
	if err := http.ListenAndServe(":8080", apiServer.Router()); err != nil {
		log.Fatal(err)
	}
}
