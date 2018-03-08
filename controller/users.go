package controller

import (
	"net/http"
	"github.com/gorilla/mux"

	"helpers"
	"app/model"
)

func GetUsers(res http.ResponseWriter, req *http.Request) {

	model := &model.UsersModel{}
	data := model.Init().SetCollection("users").GetAll()
	
	helpers.Render(res, data)
}

func GetUser(res http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	model := &model.UsersModel{}
	data := model.Init().SetCollection("users", vars["documentId"]).GetOne()

	helpers.Render(res, data)
}

func CreateUser(res http.ResponseWriter, req *http.Request) {

	body := helpers.JsonDecode(res, req)
	model := &model.UsersModel{}
	data := model.Init().Create( body )

	helpers.Render(res, data)
}

func UpdateUser(res http.ResponseWriter, req *http.Request) {

	body := helpers.JsonDecode(res, req)
	vars := mux.Vars(req)
	model := &model.UsersModel{}
	data := model.Init().Update(vars["documentId"], body)

	helpers.Render(res, data)
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	model := &model.UsersModel{}
	data := model.Init().Delete(vars["documentId"])

	helpers.Render(res, data)
}