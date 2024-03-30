package indexRoute

import (
	"github.com/gorilla/mux"
	"net/http"

	configs "GO_PROJECT/config"
	"GO_PROJECT/routes/v1"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("All OK"))
}

func RegisterRoutes(r *mux.Router) {

	subRouter := r.PathPrefix(configs.Config.APP.PREFIX + "/api/v1").Subrouter()

	subRouter.HandleFunc("/health", healthHandler).Methods("GET")
	// routes.RegisterHomeRoutes(subRouter)
	routes.RegisterSignupRoutes(subRouter)
	routes.RegisterLoginRoutes(subRouter)
	routes.RegisterHomeRoutes(subRouter)
	routes.RegisterLogoutRoutes(subRouter)
}
