package routes

import (
	"github.com/rogaha/go-postgres-jwt-react-starter/server/modules/payment/controller"
	"github.com/rogaha/go-postgres-jwt-react-starter/server/shared"
)

//GetRoutes - export routes from module
func GetRoutes() []shared.Route {
	routes := []shared.Route{
		{Method: "GET", Pattern: "/payments", HandlerFunc: controller.GetAll, UseAuth: true},
		{Method: "GET", Pattern: "/payments/:id", HandlerFunc: controller.GetByID, UseAuth: true},
		{Method: "POST", Pattern: "/payments", HandlerFunc: controller.CreateOrUpdate, UseAuth: true},
		{Method: "DELETE", Pattern: "/payments/:id", HandlerFunc: controller.Delete, UseAuth: true},
	}

	return routes
}
