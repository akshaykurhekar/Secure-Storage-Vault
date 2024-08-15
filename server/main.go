package main

import (
	"fmt"
	"net/http"
	"server/db"
	"server/routes"

	"github.com/gorilla/handlers"
)


func main() {
	
    db.Init()
    //setup router
    router := routes.SetUpRoutes()
	
    // define the port to listen on
    port := "5000"
    fmt.Printf("Starting server at port %s\n", port)

    // start the HTTP server
    if err := http.ListenAndServe(":"+port, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET","POST","DELETE","PUT","OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With","Content-Type"}),
	)(router)); err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }

}