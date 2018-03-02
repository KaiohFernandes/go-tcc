package bootstrap

import (
	"routes"
	"database"
	"log"
	"net/http"
)

func Initialize() {
	
	db := database.Database{}
	dbInit := db.Initialize()

	log.Println(dbInit)

	route := routes.Routes{}
	route.RouteHandler();

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}