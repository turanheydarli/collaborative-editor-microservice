package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turanheydarli/collaborative-editor/internal/auth-service/routes"
)

func main() {
	router := mux.NewRouter()

	routes.RegisterRoutes(router)

	log.Printf("Auth Service is running at :8086")
	log.Fatal(http.ListenAndServe(":8086", router))
}
