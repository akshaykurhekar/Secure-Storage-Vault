package routes

import (
	"github.com/gorilla/mux"
	"server/handlers"
)

func SetUpRoutes() *mux.Router {

	route := mux.NewRouter()

	route.HandleFunc("/test", handlers.TestFunc).Methods("GET")
	route.HandleFunc("/get/vaults", handlers.GetAllVaults).Methods("GET")
	route.HandleFunc("/create/vault", handlers.CreateVault).Methods("POST")
	route.HandleFunc("/update/vault/{id}", handlers.UpdateVault).Methods("PUT")
	route.HandleFunc("/delete/vault/{id}", handlers.DeleteVaultById).Methods("DELETE")

	return route
}

