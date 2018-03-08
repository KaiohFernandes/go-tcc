package routes

import (
	"app/controller"
)

func UserRoute(route *Routes){

	route.Get("/users", controller.GetUsers)
	route.Get("/users/{documentId}", controller.GetUser)
	route.Post("/users", controller.CreateUser)
	route.Put("/users/{documentId}", controller.UpdateUser)
	route.Delete("/users/{documentId}", controller.DeleteUser)
}
