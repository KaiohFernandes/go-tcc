package routes

func (route *Routes) routeProvider() {
	UserRoute(route)
	ListRoute(route)
}
