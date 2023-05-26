package routes

import (
	"github.com/aiyanhamid/mta-hosting-optimizer/app/controllers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/hostnames", controllers.GetHostnamesWithLessOrEqualActiveIPs).Methods("GET")
}
