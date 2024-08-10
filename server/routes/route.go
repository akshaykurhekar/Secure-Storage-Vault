package routes

import (
	"github.com/gorilla/mux"
	"server/handlers"
)

func SetUpRoutes() *mux.Router {

	route := mux.NewRouter()

	route.HandleFunc("/test", handlers.TestFunc).Methods("GET")
	route.HandleFunc("/getAllVaults", handlers.GetAllVaults).Methods("GET")

	return route
}

