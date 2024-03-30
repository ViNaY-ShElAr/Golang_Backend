package routes

import (
	"github.com/gorilla/mux"

	homeController "GO_PROJECT/controllers/home"
	"GO_PROJECT/middlewares/auth"
)

func RegisterHomeRoutes(r *mux.Router) {
	r.HandleFunc("/home", auth.Auth(homeController.HomeHandler)).Methods("GET")
}
