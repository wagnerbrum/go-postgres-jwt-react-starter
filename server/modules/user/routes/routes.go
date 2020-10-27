package routes

import (
	"github.com/rogaha/go-postgres-jwt-react-starter/server/modules/user/controller"
	"github.com/rogaha/go-postgres-jwt-react-starter/server/shared"
)

//GetRoutes - export routes from module
func GetRoutes() []shared.Route {
	routes := []shared.Route{
		{Method: "POST", Pattern: "/login", HandlerFunc: controller.Login, UseAuth: false},
		{Method: "POST", Pattern: "/register", HandlerFunc: controller.Create, UseAuth: false},
		{Method: "GET", Pattern: "/session", HandlerFunc: controller.Session, UseAuth: true},
	}

	return routes
}
