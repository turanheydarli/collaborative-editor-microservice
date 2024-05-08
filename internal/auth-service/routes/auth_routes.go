package routes

import (
	"github.com/gorilla/mux"
	"github.com/turanheydarli/collaborative-editor/internal/auth-service/handlers"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
}
