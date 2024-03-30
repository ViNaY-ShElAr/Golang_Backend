package routes

import (
	"github.com/gorilla/mux"

	logoutController "GO_PROJECT/controllers/logout"
	"GO_PROJECT/middlewares/auth"
)

func RegisterLogoutRoutes(r *mux.Router) {
	r.HandleFunc("/logout/user", auth.Auth(logoutController.LogoutHandler)).Methods("POST")
}
