package routes

import (
	"github.com/gorilla/mux"
	"server/handlers"
)

func SetUpRoutes() *mux.Router {

	route := mux.NewRouter()

	route.HandleFunc("/test", handlers.TestFunc).Methods("GET")
	route.HandleFunc("/get/vaults", handlers.GetAllVaults).Methods("GET")
	route.HandleFunc("/get/credentials/{vid}", handlers.GetCredentialByVaultId).Methods("GET")
	route.HandleFunc("/create/credentials", handlers.CreateCredential).Methods("POST")
	route.HandleFunc("/create/vault", handlers.CreateVault).Methods("POST")
	route.HandleFunc("/update/vault/{id}", handlers.UpdateVault).Methods("PUT")
	route.HandleFunc("/update/credential/{id}", handlers.UpdateCredential).Methods("PUT")
	route.HandleFunc("/delete/vault/{id}", handlers.DeleteVaultById).Methods("DELETE")
	route.HandleFunc("/delete/credential/{id}", handlers.DeleteCredential).Methods("DELETE")

	return route
}

