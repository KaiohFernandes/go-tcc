package controller

import (
	"fmt"
	"net/http"
	"log"

	"app/model"
)

func Index(req http.ResponseWriter, res *http.Request) {

	db := model.Init()

	for _, d := range db.GetUsers(){
		log.Println(d["first"])
   	}
	
	fmt.Fprintf(req, "Inicio")
}

func Sobre(req http.ResponseWriter, res *http.Request) {
	fmt.Fprintf(req, "Sobre")
}

func Contato(req http.ResponseWriter, res *http.Request) {
	fmt.Fprintf(req, "Contato")
}