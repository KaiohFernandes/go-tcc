package controller

import (
	"net/http"
	"github.com/gorilla/mux"

	"helpers"
	"app/model"
)

func GetUsers(res http.ResponseWriter, req *http.Request) {

	db := model.UsersInit()
	data := db.GetUsers()
	
	helpers.Render(res, data)
}

func GetUser(res http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	data := model.UsersInit().GetUser( vars["documentId"] )

	helpers.Render(res, data)
}

func CreateUser(res http.ResponseWriter, req *http.Request) {

	body := helpers.JsonDecode(res, req)
	data := model.UsersInit().CreateUser( body )

	helpers.Render(res, data)
}

func UpdateUser(res http.ResponseWriter, req *http.Request) {

	body := helpers.JsonDecode(res, req)
	vars := mux.Vars(req)
	data := model.UsersInit().UpdateUser(vars["documentId"], body)

	helpers.Render(res, data)
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	data := model.UsersInit().DeleteUser(vars["documentId"])

	helpers.Render(res, data)
}