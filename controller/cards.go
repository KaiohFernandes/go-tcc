package controller

import (
	"net/http"
	"github.com/gorilla/mux"

	"helpers"
	"app/model"
)

func GetCardList(res http.ResponseWriter, req *http.Request) {

	model := &model.Model{}
	vars := mux.Vars(req)

	data := model.Init().SetCollection("cards").GetQuery("listId", "==", vars["listId"] )
	helpers.Render(res, data);

}

func GetCards(res http.ResponseWriter, req *http.Request) {

	model := &model.Model{}
	data := model.Init().SetCollection("cards").GetAll()
	
	helpers.Render(res, data)
}

func GetCard(res http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	model := &model.Model{}
	data := model.Init().SetCollection("cards", vars["documentId"]).GetOne()

	helpers.Render(res, data)
}

func CreateCard(res http.ResponseWriter, req *http.Request) {

	body := helpers.JsonDecode(res, req)
	model := &model.Model{}
	data := model.Init().SetCollection("cards").Create( body )

	helpers.Render(res, data)
}

func UpdateCard(res http.ResponseWriter, req *http.Request) {

	body := helpers.JsonDecode(res, req)
	vars := mux.Vars(req)
	model := &model.Model{}
	data := model.Init().SetCollection("cards", vars["documentId"]).Update(body)

	helpers.Render(res, data)
}

func DeleteCard(res http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	model := &model.Model{}
	data := model.Init().SetCollection("cards", vars["documentId"]).Delete()

	helpers.Render(res, data)
}