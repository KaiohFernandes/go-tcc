package controller

import (
	"net/http"
	"github.com/gorilla/mux"

	"helpers"
	"app/model"
	"fmt"
	"log"
)

func GetToken(res http.ResponseWriter, req *http.Request) {

	model := &model.Model{}
	body := helpers.JsonDecode(res, req);

	if body["email"] != nil { 
		data := model.Init().SetCollection("users").GetQuery("email", "==", body["email"].(string))

		if data[0]["password"].(string) == body["password"].(string) {
			
			auth := make(map[string] interface{}) 
			auth["email"] = data[0]["email"]
			auth["username"] = data[0]["username"]
			auth["id"] = data[0]["id"]

			log.Println( auth )

			helpers.Render(res, auth);
		} else {
			fmt.Fprintf(res, "{'error': 'Invalid email or password'}")
		}

	} else {
		fmt.Fprintf(res, "{'error':'Invalid values'}")
	}

}

func GetUsers(res http.ResponseWriter, req *http.Request) {

	model := &model.Model{}
	data := model.Init().SetCollection("users").GetAll()
	helpers.Render(res, data)
}

func GetUser(res http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	model := &model.Model{}
	data := model.Init().SetCollection("users", vars["documentId"]).GetOne()

	helpers.Render(res, data)
}

func CreateUser(res http.ResponseWriter, req *http.Request) {

	body := helpers.JsonDecode(res, req)
	model := &model.Model{}
	data := model.Init().SetCollection("users").Create( body )

	helpers.Render(res, data)
}

func UpdateUser(res http.ResponseWriter, req *http.Request) {

	body := helpers.JsonDecode(res, req)
	vars := mux.Vars(req)
	model := &model.Model{}
	data := model.Init().SetCollection("users", vars["documentId"]).Update(body)

	helpers.Render(res, data)
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	model := &model.Model{}
	data := model.Init().SetCollection("users", vars["documentId"]).Delete()

	helpers.Render(res, data)
}