package routes

import (
	"app/controller"
)

func ListRoute(route *Routes){

	route.Get("/lists", controller.GetLists)
	route.Get("/lists/{documentId}", controller.GetList)
	route.Post("/lists", controller.CreateList)
	route.Get("/lists/users/{userId}", controller.GetListUser)
	route.Put("/lists/{documentId}", controller.UpdateList)
	route.Delete("/lists/{documentId}", controller.DeleteList)
}
