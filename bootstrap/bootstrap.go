package bootstrap

import (
	"app/routes"
	"log"
	"net/http"
)

func Initialize() {

	route := &routes.Routes{}
	route.RouteHandler();

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}