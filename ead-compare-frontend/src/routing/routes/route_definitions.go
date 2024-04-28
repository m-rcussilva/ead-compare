package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                   string
	Method                string
	HandlerFunc           func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

func ConfigureRouter(router *mux.Router) *mux.Router {
	r := publicRoutes
	r = append(r, registerRoutes...)
	r = append(r, registeredRoutes...)

	for _, route := range r {
		router.HandleFunc(route.URI, route.HandlerFunc).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
