package server

import (
	"net/http"

	"github.com/earlgray283/material-gakujo/api/db"
	auth "github.com/earlgray283/material-gakujo/api/server/libauth"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type ApiServer struct {
	controller        *db.Controller
	cryptoKey         []byte
	sessionController *auth.SessionController
}

func NewApiServer(config *db.DBConfig, cryptoKey []byte) (*ApiServer, error) {
	controller, err := db.NewController(config)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	sessionController := auth.NewSessionController(controller)

	return &ApiServer{
		controller,
		cryptoKey,
		sessionController,
	}, nil
}

func (api *ApiServer) Router() *mux.Router {
	r := mux.NewRouter()

	r.Use(LoggingMiddleware, AuthMiddleware(api.sessionController))
	r.HandleFunc("/api/auth/login", api.Login).Methods(http.MethodPost).Headers("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	r.HandleFunc("/api/auth/logout", api.Logout).Methods(http.MethodPost)
	r.HandleFunc("/api/auth/register", api.RegistNewUser).Methods(http.MethodPost).Headers("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	r.HandleFunc("/api/crawle", api.Crawle).Methods(http.MethodPut)
	r.HandleFunc("/api/seisekis", api.FetchSeiseki).Methods(http.MethodGet)
	r.HandleFunc("/api/profile", api.FetchSeiseki).Methods(http.MethodGet)

	return r
}
