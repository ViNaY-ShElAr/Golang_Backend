package routes

import (
	"github.com/gorilla/mux"

	signupController "GO_PROJECT/controllers/signup"
)

func RegisterSignupRoutes(r *mux.Router) {
	r.HandleFunc("/signup/user", signupController.SignupHandler).Methods("POST")
}
