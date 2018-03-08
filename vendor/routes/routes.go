package routes

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Routes struct{
	Router *mux.Router
}

func (route *Routes) RouteHandler() {
	route.Router = mux.NewRouter()

	route.routeProvider()

	http.Handle("/", route.Router)
}

// Request Verbs

func (route *Routes) Get(path string, routeFunc func(http.ResponseWriter, *http.Request)) {
	route.Router.HandleFunc(path, routeFunc).Methods("GET")
}

func (route *Routes) Post(path string, routeFunc func(http.ResponseWriter, *http.Request)) {
	route.Router.HandleFunc(path, routeFunc).Methods("POST")
}

func (route *Routes) Put(path string, routeFunc func(http.ResponseWriter, *http.Request)) {
	route.Router.HandleFunc(path, routeFunc).Methods("PUT")
}

func (route *Routes) Delete(path string, routeFunc func(http.ResponseWriter, *http.Request)) {
	route.Router.HandleFunc(path, routeFunc).Methods("DELETE")
}