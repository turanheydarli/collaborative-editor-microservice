package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turanheydarli/collaborative-editor/services/auth-service/internal/routes"
	"github.com/turanheydarli/collaborative-editor/services/auth-service/internal/services"
)

func main() {
	router := mux.NewRouter()

	services.InitializeStorage()

	routes.RegisterRoutes(router)

	log.Printf("Auth Service is running at :8086")
	log.Fatal(http.ListenAndServe(":8086", router))
}
