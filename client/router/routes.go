package router

import (
	"client/controller"

	"github.com/labstack/echo/v4"
)

func Echo(e *echo.Echo, uc controller.AuthController, oc controller.OrderController) {
	users := e.Group("/users")
	{
		users.POST("/login", uc.LoginUser)
		users.POST("/register", uc.RegisterUser)
	}

	orders := e.Group("/orders")
	{
		orders.POST("/create", oc.CreateOrder)
		orders.GET("/", oc.GetAllOrders)
		orders.DELETE("/", oc.DeleteOrder)
	}

}
