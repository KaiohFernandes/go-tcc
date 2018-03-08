package controller

import (
	"fmt"
	"net/http"
	_ "log"
	"encoding/json"

	"app/model"
)

func Index(req http.ResponseWriter, res *http.Request) {

	db := model.Init()
	data := db.GetUsers()
	jsonString, _ := json.Marshal(data)

	fmt.Fprintf(req, string(jsonString))
}

func Sobre(req http.ResponseWriter, res *http.Request) {
	fmt.Fprintf(req, "Sobre")
}

func Contato(req http.ResponseWriter, res *http.Request) {
	fmt.Fprintf(req, "Contato")
}