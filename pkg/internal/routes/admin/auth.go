package admin

import (
	"internal/controllers"

	"github.com/gorilla/mux"
)

func addAuth(router *mux.Router) {
	auth := router.PathPrefix("/auth").Subrouter()

	auth.HandleFunc("/login", controllers.AuthLogin).Methods("POST")
	auth.HandleFunc("/logout", controllers.AuthLogout).Methods("POST")
	auth.HandleFunc("/forgot-password", controllers.AuthForgotPassword).Methods("POST")
}
