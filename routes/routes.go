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

// Get all routes
func (route *Routes) routeProvider() {
	UserRoute(route)
}

func (route *Routes) Get(path string, routeFunc func(http.ResponseWriter, *http.Request)) {
	route.Router.HandleFunc(path, routeFunc).Methods("GET")
}

