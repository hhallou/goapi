package healthz

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes ...
func Routes(router *mux.Router, handler Handler) http.Handler {
	healthzRouter := router.PathPrefix("/").Subrouter().StrictSlash(true)
	healthzRouter.HandleFunc("/healthz", handler.getHealthz)

	return healthzRouter
}
