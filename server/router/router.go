package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rogaha/go-postgres-jwt-react-starter/server/middlewares"
	customer_routes "github.com/rogaha/go-postgres-jwt-react-starter/server/modules/customer/routes"
	payment_routes "github.com/rogaha/go-postgres-jwt-react-starter/server/modules/payment/routes"
	user_routes "github.com/rogaha/go-postgres-jwt-react-starter/server/modules/user/routes"
	"github.com/rogaha/go-postgres-jwt-react-starter/server/shared"
)

// SetupRouter setup routing here
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Middlewares
	router.Use(middlewares.ErrorHandler)
	router.Use(middlewares.CORSMiddleware())

	// Route Test API OK
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	// Registry Routes
	listUserRoutes := user_routes.GetRoutes()
	listCustomerRoutes := customer_routes.GetRoutes()
	listPaymentRoutes := payment_routes.GetRoutes()

	registerRoutes(router, listUserRoutes)
	registerRoutes(router, listCustomerRoutes)
	registerRoutes(router, listPaymentRoutes)

	return router
}

func registerRoutes(router *gin.Engine, routes []shared.Route) {
	for _, route := range routes {
		registerRoute(router, route)
	}
}

func registerRoute(router *gin.Engine, route shared.Route) {
	switch route.Method {
	case "GET":
		router.GET(route.Pattern, middlewares.AuthMiddleware(route.UseAuth), route.HandlerFunc)
	case "POST":
		router.POST(route.Pattern, middlewares.AuthMiddleware(route.UseAuth), route.HandlerFunc)
	case "PUT":
		router.PUT(route.Pattern, middlewares.AuthMiddleware(route.UseAuth), route.HandlerFunc)
	case "DELETE":
		router.DELETE(route.Pattern, middlewares.AuthMiddleware(route.UseAuth), route.HandlerFunc)
	default:
		router.GET(route.Pattern, func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"result": "Specify a valid http method with this route.",
			})
		})
	}
}
