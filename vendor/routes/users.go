package routes

import (
	"app/controller"
)

func UserRoute(route *Routes){

	route.Get("/users", controller.Index)

}
