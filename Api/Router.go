package api

import (
	"avitoTest/api/handlers/ping_handler"

	"github.com/gorilla/mux"
)

// InitRoutes - initializes API routes
func InitRoutes(router *mux.Router) {
	router.HandleFunc("/api/ping", ping_handler.PingHandler).Methods("GET")
}
