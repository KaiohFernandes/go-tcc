package bootstrap

import (
	"routes"
	"database"
	"log"
	"net/http"
)

func Initialize() {
	
	db := database.Firebase{}
	db.Initialize()

	route := routes.Routes{}
	route.RouteHandler();

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3001", nil))
}