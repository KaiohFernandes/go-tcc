package controller

import (
	"net/http"
	"github.com/gorilla/mux"

	"helpers"
	"app/model"
)

func GetCardStyle(res http.ResponseWriter, req *http.Request) {

	model := &model.Model{}
	vars := mux.Vars(req)
	data := model.Init().SetCollection("cards", vars["documentId"], "styles").GetAll()
	
	helpers.Render(res, data)
}

func CreateCardStyle(res http.ResponseWriter, req *http.Request) {

	body := helpers.JsonDecode(res, req)
	model := &model.Model{}
	vars := mux.Vars(req)
	data := model.Init().SetCollection("cards", vars["documentId"], "styles").Create( body )

	helpers.Render(res, data)
}

func UpdateCardStyle(res http.ResponseWriter, req *http.Request) {

	body := helpers.JsonDecode(res, req)
	vars := mux.Vars(req)
	model := &model.Model{}
	data := model.Init().SetCollection("cards", vars["documentId"], "styles", vars["styleId"]).Update(body)

	helpers.Render(res, data)
}

func DeleteCardStyle(res http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	model := &model.Model{}
	data := model.Init().SetCollection("cards", vars["documentId"], "styles", vars["styleId"]).Delete()

	helpers.Render(res, data)
}