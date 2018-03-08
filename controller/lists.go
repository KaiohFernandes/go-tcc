package controller

import (
	"net/http"
	"github.com/gorilla/mux"

	"helpers"
	"app/model"
)

func GetLists(res http.ResponseWriter, req *http.Request) {

	model := &model.ListsModel{}
	data := model.Init().SetCollection("lists").GetAll()
	
	helpers.Render(res, data)
}

func GetList(res http.ResponseWriter, req *http.Request) {
	
	vars := mux.Vars(req)

	model := &model.ListsModel{}
	data := model.Init().SetCollection("lists", vars["documentId"]).GetOne()

	helpers.Render(res, data)
}

func CreateList(res http.ResponseWriter, req *http.Request) {

	body := helpers.JsonDecode(res, req)

	model := &model.ListsModel{}
	data := model.Init().Create( body )

	helpers.Render(res, data)
}

func UpdateList(res http.ResponseWriter, req *http.Request) {

	body := helpers.JsonDecode(res, req)
	vars := mux.Vars(req)

	model := &model.ListsModel{}
	data := model.Init().Update(vars["documentId"], body)

	helpers.Render(res, data)
}

func DeleteList(res http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)

	model := &model.ListsModel{}
	data := model.Init().Delete(vars["documentId"])

	helpers.Render(res, data)
}