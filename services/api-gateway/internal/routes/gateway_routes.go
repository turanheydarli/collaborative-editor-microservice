package routes

import (
	"github.com/gorilla/mux"
	"github.com/turanheydarli/collaborative-editor/services/api-gateway/internal/handlers"
)

func RegisterAuthRoutes(router *mux.Router) {
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
}
