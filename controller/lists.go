package controller

import (
	"net/http"
	"github.com/gorilla/mux"

	"helpers"
	"app/model"
)

func GetListUser(res http.ResponseWriter, req *http.Request) {

	model := &model.Model{}
	vars := mux.Vars(req)

	data := model.Init().SetCollection("lists").GetQuery("userId", "==", vars["userId"] )
	helpers.Render(res, data);

}

func GetLists(res http.ResponseWriter, req *http.Request) {

	model := &model.Model{}
	data := model.Init().SetCollection("lists").GetAll()
	
	helpers.Render(res, data)
}

func GetList(res http.ResponseWriter, req *http.Request) {
	
	vars := mux.Vars(req)

	model := &model.Model{}
	data := model.Init().SetCollection("lists", vars["documentId"]).GetOne()

	helpers.Render(res, data)
}

func CreateList(res http.ResponseWriter, req *http.Request) {

	body := helpers.JsonDecode(res, req)

	model := &model.Model{}
	data := model.Init().SetCollection("lists").Create( body )

	helpers.Render(res, data)
}

func UpdateList(res http.ResponseWriter, req *http.Request) {

	body := helpers.JsonDecode(res, req)
	vars := mux.Vars(req)

	model := &model.Model{}
	data := model.Init().SetCollection("lists", vars["documentId"]).Update(body)

	helpers.Render(res, data)
}

func DeleteList(res http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)

	model := &model.Model{}
	data := model.Init().SetCollection("lists", vars["documentId"]).Delete()

	helpers.Render(res, data)
}