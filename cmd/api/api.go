package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/chandruchiku/go-ecom/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func New(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter)

	log.Println("Starting server on", s.addr)
	error := http.ListenAndServe(s.addr, nil)
	if error != nil {
		return error
	}
	return nil
}
