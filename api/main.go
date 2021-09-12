package main

import (
	"log"
	"net/http"
	"os"

	"github.com/earlgray283/material-gakujo/api/db"
	"github.com/earlgray283/material-gakujo/api/server"
	"github.com/gorilla/handlers"
	"github.com/rs/cors"
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

	router := apiServer.Router()
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{http.MethodPost, http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler(router)
	loggerRouter := handlers.LoggingHandler(os.Stdout, corsHandler)

	log.Println("Server Started")
	if err := http.ListenAndServe(":8080", loggerRouter); err != nil {
		log.Fatal(err)
	}
}
