package routing

import "github.com/gorilla/mux"

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	return r
}
