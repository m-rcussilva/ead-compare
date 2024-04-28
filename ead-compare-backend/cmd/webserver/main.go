package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	// "github.com/gorilla/mux"
	"github.com/m-rcussilva/go-private/tree/main/2024/projects/02-ead-compare/ead-compare-backend/src/config"
	"github.com/m-rcussilva/go-private/tree/main/2024/projects/02-ead-compare/ead-compare-backend/src/config/logger"
	"github.com/m-rcussilva/go-private/tree/main/2024/projects/02-ead-compare/ead-compare-backend/src/routing"
)

func main() {
	logger.InitLoger()
	config.LoadDatabaseEnv()

	router := routing.CreateRouter()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	routerWithCORS := handlers.CORS(originsOk, headersOk, methodsOk)(router)

	log.Println("Starting API")
	fmt.Println("Server is up and running on port: http://localhost:5000")

	log.Fatal(http.ListenAndServe(":5000", routerWithCORS))
}
