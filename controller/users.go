package controller

import (
	"fmt"
	"net/http"
	"encoding/json"

	"app/model"
)

func Index(req http.ResponseWriter, res *http.Request) {

	db := model.Init()
	users := db.GetUsers()

	jsonString, _ := json.Marshal(users)
	
	fmt.Fprintf(req, string( jsonString ) )
}

func Sobre(req http.ResponseWriter, res *http.Request) {
	fmt.Fprintf(req, "Sobre")
}

func Contato(req http.ResponseWriter, res *http.Request) {
	fmt.Fprintf(req, "Contato")
}