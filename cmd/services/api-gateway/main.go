package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turanheydarli/collaborative-editor/internal/api-gateway/routes"
)

func main() {
	router := mux.NewRouter()

	routes.RegisterAuthRoutes(router)

	log.Printf("API Gateway is running at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
