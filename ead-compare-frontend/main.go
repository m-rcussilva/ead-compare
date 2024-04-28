package main

import (
	"fmt"
	"frontend/src/routing"
	"frontend/src/routing/routes"
	"frontend/src/utils"
	"log"
	"net/http"
)

func main() {
	utils.LoadTemplates()

	r := routing.SetupRouter()
	r = routes.ConfigureRouter(r)

	fmt.Println("http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
