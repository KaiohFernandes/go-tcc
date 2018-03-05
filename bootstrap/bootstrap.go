package bootstrap

import (
	"routes"
	"database"
	"log"
	"fmt"
	"net/http"
)

func Initialize() {
	
	db := database.Init();
	client := db.OpenConnection();
	data := db.Get(client, "users")
	fmt.Println(data)
	defer db.CloseConnection(client)

	route := routes.Routes{}
	route.RouteHandler();

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}