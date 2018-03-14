package routes

import (
	"app/controller"
)

func CardsRoute(route *Routes){

	route.Get("/cards", controller.GetCards)
	route.Get("/cards/{documentId}", controller.GetCard)
	route.Post("/cards", controller.CreateCard)
	route.Put("/cards/{documentId}", controller.UpdateCard)
	route.Delete("/cards/{documentId}", controller.DeleteCard)
}
