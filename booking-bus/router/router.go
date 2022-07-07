package router

import (
	"github.com/archit-tiket/booking-system/booking-bus/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/buses", middleware.GetAllBuses).Methods("GET", "OPTIONS")

	return router
}
