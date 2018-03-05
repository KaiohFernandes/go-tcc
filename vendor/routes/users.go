package routes

import (
	"app/controller"
)

func UserRoute(route *Routes){

	route.Get("/", controller.Index)
	route.Get("/sobre", controller.Sobre)
	route.Get("/contato", controller.Contato)

}
