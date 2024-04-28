package routing

import (
	"github.com/gorilla/mux"
	"github.com/m-rcussilva/go-private/tree/main/2024/projects/02-ead-compare/ead-compare-backend/src/routing/routes"
)

func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigureRouter(r)
}
