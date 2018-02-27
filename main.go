package main

import (
	"app/routes"
	"log"
	"net/http"
)

func main() {

	rotas := routes.Routes{}

	http.HandleFunc("/", rotas.RouteHandler)

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
