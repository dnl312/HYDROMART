package router

import (
	"client/controller"

	"github.com/labstack/echo/v4"
)

func Echo(e *echo.Echo, uc controller.AuthController, mc controller.MerchantController, oc controller.OrderController) {
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

	merchants := e.Group("/merchants")
	{
		merchants.GET("/products", mc.ShowAllProducts)
		merchants.POST("/products", mc.AddProduct)
		merchants.PUT("/products/:product_id", mc.UpdateProduct)
		merchants.DELETE("/products/:product_id", mc.DeleteProduct)
		merchants.GET("/orders", mc.ShowAllOrders)
		merchants.PUT("/orders/:order_id", mc.ProcessOrder)
	}

}
