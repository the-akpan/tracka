package admin

import "github.com/gorilla/mux"

func Register(router *mux.Router) {
	admin := router.PathPrefix("/admin").Subrouter()

	addAuth(admin)
	addOrders(admin)
}
