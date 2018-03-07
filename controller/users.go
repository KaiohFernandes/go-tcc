package controller

import (
	"net/http"
	"github.com/gorilla/mux"

	"helpers"
	"app/model"
)

type userFormat struct {

}

func Index(res http.ResponseWriter, req *http.Request) {

	db := model.Init()
	data := db.GetUsers()
	
	helpers.Render(res, data)
}

func Show(res http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	data := model.Init().GetUser( vars["documentId"] )

	helpers.Render(res, data)
}

func Create(res http.ResponseWriter, req *http.Request) {

	body := helpers.JsonDecode(res, req)
	data := model.Init().Create( body )

	helpers.Render(res, data)
}

func Update(res http.ResponseWriter, req *http.Request) {

	body := helpers.JsonDecode(res, req)

	vars := mux.Vars(req)
	data := model.Init().Update(vars["documentId"], body)

	helpers.Render(res, data)
}

func Delete(res http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	data := model.Init().Delete(vars["documentId"])

	helpers.Render(res, data)
}