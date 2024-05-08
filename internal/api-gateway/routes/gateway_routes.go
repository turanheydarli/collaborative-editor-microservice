package routes

import (
	"github.com/gorilla/mux"
	"github.com/turanheydarli/collaborative-editor/internal/api-gateway/handlers"
)

func RegisterAuthRoutes(router *mux.Router) {
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
}
