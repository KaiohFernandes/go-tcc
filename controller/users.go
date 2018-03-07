package controller

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"

	"app/model"
)

type userFormat struct {

}

func Index(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	db := model.Init()
	users := db.GetUsers()

	jsonString, err := json.Marshal(users)

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Fprintf(res, string( jsonString ))
}

func Show(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)
	data := model.Init().GetUser( vars["documentId"] )

	jsonString, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(res, string( jsonString ))
}

func Create(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	decode := json.NewDecoder(req.Body)
	body := make( map[string] interface{} )

	if err := decode.Decode(&body); err != nil{
		panic(err)
	}

	defer req.Body.Close()

	data := model.Init().Create( body )

	jsonString, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(res, string(jsonString))
}

func Update(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)
	decode := json.NewDecoder(req.Body)

	body := make( map[string] interface{} )

	if err := decode.Decode(&body); err != nil{
		panic(err)
	}

	defer req.Body.Close()

	data := model.Init().Update(vars["documentId"], body)

	jsonString, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(res, string(jsonString))
}

func Delete(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)

	data := model.Init().Delete(vars["documentId"])

	jsonString, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(res, string(jsonString))

}