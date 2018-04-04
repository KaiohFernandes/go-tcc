package routes

import (
	"app/controller"
)

func CardsRoute(route *Routes){

	route.Get("/cards", controller.GetCards)
	route.Get("/cards/{documentId}", controller.GetCard)
	route.Post("/cards", controller.CreateCard)
	route.Get("/cards/list/{listId}", controller.GetCardList)
	route.Put("/cards/{documentId}", controller.UpdateCard)
	route.Delete("/cards/{documentId}", controller.DeleteCard)

	// Cards Style
	route.Get("/cards/{documentId}/styles", controller.GetCardStyle)
	route.Post("/cards/{documentId}/styles", controller.CreateCardStyle)
	route.Put("/cards/{documentId}/styles/{styleId}", controller.UpdateCardStyle)
	route.Delete("/cards/{documentId}/styles/{styleId}", controller.DeleteCardStyle)

}
