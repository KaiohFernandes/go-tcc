package routes

import (
	"app/controller"
)

func UserRoute(route *Routes){

	route.Get("/users", controller.Index)
	route.Get("/users/{documentId}", controller.Show)
	route.Post("/users", controller.Create)
	route.Put("/users/{documentId}", controller.Update)
	route.Delete("/users/{documentId}", controller.Delete)
}
