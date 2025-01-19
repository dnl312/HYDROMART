package router

import (
	"client/controller"

	"github.com/labstack/echo/v4"
)

func Echo(e *echo.Echo, uc controller.AuthController, mc controller.MerchantController) {
	users := e.Group("/users")
	{
		users.POST("/login", uc.LoginUser)
		users.POST("/register", uc.RegisterUser)
	}

	merchants := e.Group("/merchants")
	{
		merchants.GET("/products", mc.ShowAllProduct)
		merchants.POST("/products", mc.AddProduct)
		// merchants.PUT("/products/:id", mc.UpdateProduct)
		// merchants.DELETE("/products/:id", mc.UpdateProduct)
		// merchants.GET("/transactions", mc.ShowTransactions)
	}

}
