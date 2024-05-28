package main

import (
	"log"
	"net/http"

	"sipinter-api/routes"
	"sipinter-api/utils"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	utils.ConnectDB()

	routes.RegisterRoutes(router)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
