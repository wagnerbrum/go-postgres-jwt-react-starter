package routes

import (
	"github.com/rogaha/go-postgres-jwt-react-starter/server/modules/customer/controller"
	"github.com/rogaha/go-postgres-jwt-react-starter/server/shared"
)

//GetRoutes - export routes from module
func GetRoutes() []shared.Route {
	routes := []shared.Route{
		{Method: "GET", Pattern: "/customers", HandlerFunc: controller.GetAll, UseAuth: true},
		{Method: "GET", Pattern: "/customers/:id", HandlerFunc: controller.GetByID, UseAuth: true},
		{Method: "GET", Pattern: "/customers/:id/paymentmethods", HandlerFunc: controller.GetAllByCustomer, UseAuth: true},
		{Method: "GET", Pattern: "/customers/:id/paymentmethods/count", HandlerFunc: controller.GetCountAllByCustomer, UseAuth: true},
		{Method: "POST", Pattern: "/customers", HandlerFunc: controller.CreateOrUpdate, UseAuth: true},
		{Method: "DELETE", Pattern: "/customers/:id", HandlerFunc: controller.Delete, UseAuth: true},
	}

	return routes
}
