package main

import (
	"fmt"
	"net/http"
    "server/routes"
)


func main() {
	

    //setup router

    router := routes.SetUpRoutes()
	
    // define the port to listen on
    port := "4000"
    fmt.Printf("Starting server at port %s\n", port)

    // start the HTTP server
    if err := http.ListenAndServe(":"+port, router); err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }

}