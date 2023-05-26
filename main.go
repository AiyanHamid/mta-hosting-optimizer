package main

import (
	"log"
	"net/http"

	"github.com/aiyanhamid/mta-hosting-optimizer/app/routes"
	"github.com/aiyanhamid/mta-hosting-optimizer/config"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	config.LoadConfig()

	port := config.GetPort()
	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
