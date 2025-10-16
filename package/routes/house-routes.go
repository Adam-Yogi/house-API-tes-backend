package routes

import (
	"gin-api/package/controllers"

	"github.com/gorilla/mux"
)

var HouseListing = func(router *mux.Router) {
	router.HandleFunc("/house/", controllers.CreateHouse).Methods("POST")
	router.HandleFunc("/house/", controllers.GetHouse).Methods("GET")
	router.HandleFunc("/house/{houseId}", controllers.GetHouseById).Methods("GET")
	router.HandleFunc("/house/{houseId}", controllers.UpdateHouse).Methods("PUT")
	router.HandleFunc("/house/{houseId}", controllers.DeleteHouse).Methods("DELETE")
}
