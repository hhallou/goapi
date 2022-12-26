package lambda

import (
	"net/http"

	"github.com/ProovGroup/claim/internal/config"
	"github.com/ProovGroup/claim/internal/healthz"
	"github.com/gorilla/mux"
)

func Serve() error {
	mainRouter := mux.NewRouter()
	lambdaRouter := newRouter(mainRouter)

	return http.ListenAndServe(":"+config.ConfigValues.AppPort, lambdaRouter)
}

func newRouter(mainRouter *mux.Router) http.Handler {
	router := mainRouter.PathPrefix("/").Subrouter().StrictSlash(true)

	healthz.Routes(router, healthz.New())

	return router
}
