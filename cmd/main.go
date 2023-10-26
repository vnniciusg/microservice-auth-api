package main

import (
	"net/http"

	"github.com/vnniciusg/microservice-auth-api/adapters/http/rest/routes"
)

func main() {
	router := routes.InitRoutes()

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
