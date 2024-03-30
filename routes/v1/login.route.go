package routes

import (
	"github.com/gorilla/mux"

	loginController "GO_PROJECT/controllers/login"
)

func RegisterLoginRoutes(r *mux.Router) {
	r.HandleFunc("/login/user", loginController.LoginHandler).Methods("POST")
}
