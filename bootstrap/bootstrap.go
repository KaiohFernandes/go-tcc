package bootstrap

import (
	"routes"
	"database"
	"log"
	"fmt"
	"net/http"
)

func Initialize() {
	
	db := database.Firebase{}
	client := db.Initialize()
	data := db.Get(client, "users")
	fmt.Println(data)
	db.CloseConnection(client)

	route := routes.Routes{}
	route.RouteHandler();

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3001", nil))
}