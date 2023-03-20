package admin

import (
	"internal/controllers"

	"github.com/gorilla/mux"
)

func addOrders(router *mux.Router) {
	orders := router.PathPrefix("/orders").Subrouter()

	orders.HandleFunc("", controllers.OrdersList).Methods("GET")
	orders.HandleFunc("", controllers.OrdersCreate).Methods("POST")
	orders.HandleFunc("/{id}", controllers.OrdersDetail).Methods("GET")
	orders.HandleFunc("/{id}", controllers.OrdersDelete).Methods("DELETE")
	orders.HandleFunc("/{id}", controllers.OrdersUpdate).Methods("PUT")
}
