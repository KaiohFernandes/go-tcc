package controller

import (
	"net/http"
	"github.com/gorilla/mux"

	"helpers"
	"app/model"
)

func GetLists(res http.ResponseWriter, req *http.Request) {

	data := model.ListsInit().GetLists()
	
	helpers.Render(res, data)
}

func GetList(res http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	data := model.ListsInit().GetList( vars["documentId"] )

	helpers.Render(res, data)
}

func CreateList(res http.ResponseWriter, req *http.Request) {

	body := helpers.JsonDecode(res, req)
	data := model.ListsInit().CreateList( body )

	helpers.Render(res, data)
}

func UpdateList(res http.ResponseWriter, req *http.Request) {

	body := helpers.JsonDecode(res, req)
	vars := mux.Vars(req)
	data := model.ListsInit().UpdateList(vars["documentId"], body)

	helpers.Render(res, data)
}

func DeleteList(res http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	data := model.ListsInit().DeleteList(vars["documentId"])

	helpers.Render(res, data)
}