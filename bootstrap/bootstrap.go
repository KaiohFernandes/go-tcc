package bootstrap

import (
	"routes"
	"log"
	"net/http"
)

func Initialize() {

	route := routes.Routes{}
	route.RouteHandler();

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3002", nil))
}