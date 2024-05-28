package routes

import (
	"sipinter-api/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/services", handlers.GetServices).Methods("GET")
	router.HandleFunc("/api/services/{id}", handlers.GetService).Methods("GET")
	router.HandleFunc("/api/services", handlers.CreateService).Methods("POST")
	router.HandleFunc("/api/services/{id}", handlers.UpdateService).Methods("PUT")
	router.HandleFunc("/api/services/{id}", handlers.DeleteService).Methods("DELETE")
	router.HandleFunc("/api/users/register", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/api/users/login", handlers.LoginUser).Methods("POST")
}
