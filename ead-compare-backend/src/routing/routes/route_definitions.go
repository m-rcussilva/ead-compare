package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RouteDefinitions struct {
	URI                   string
	Method                string
	HandlerFunc           func(w http.ResponseWriter, r *http.Request)
	RequireAuthentication bool
}

func ConfigureRouter(r *mux.Router) *mux.Router {
	appRoutes := append(userRoutes, loginRoute...)

	for _, route := range appRoutes {
		r.HandleFunc(route.URI, route.HandlerFunc).Methods(route.Method)
		log.Printf("Rota %s configurada com sucesso", route.URI)
	}

	return r
}
