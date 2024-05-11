package routes

import (
	"github.com/gorilla/mux"
	"github.com/turanheydarli/collaborative-editor/services/auth-service/internal/handlers"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
}
