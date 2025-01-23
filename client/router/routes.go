package router

import (
	"client/controller"
	"client/middleware"

	"github.com/labstack/echo/v4"
)

func Echo(e *echo.Echo, uc controller.AuthController, mc controller.MerchantController, oc controller.OrderController) {
	users := e.Group("/users")
	{
		users.POST("/login", uc.LoginUser)
		users.POST("/register", uc.RegisterUser)
		users.POST("/register/merchant", uc.RegisterMerchant)
	}

	orders := e.Group("/orders")
	orders.Use(middleware.RequireAuth)
	{
		orders.POST("/create", oc.CreateOrder)
		orders.GET("/order-list", oc.GetAllOrders)
		orders.PUT("/order-update", oc.UpdateOrder)
		orders.DELETE("/order-delete", oc.DeleteOrder)
		orders.POST("/topup", oc.TopUp)
	}

	// e.GET("/orders/update-deposit", oc.UpdateDepositCron)
	// c := cron.New()
	// _, err := c.AddFunc("@every 5s", func() {
	// 	_, err := http.Get("http://localhost:8080/orders/update-deposit")
	// 	if err != nil {
	// 		fmt.Println("Error triggering update-deposit endpoint:", err)
	// 	}
	// })
	// if err != nil {
	// 	fmt.Println("Error setting up cron job:", err)
	// }
	// c.Start()

	merchants := e.Group("/merchants")
	merchants.Use(middleware.RequireAuth)
	{
		merchants.GET("/products", mc.ShowAllProducts)
		merchants.POST("/products", mc.AddProduct)
		merchants.PUT("/products/:product_id", mc.UpdateProduct)
		merchants.DELETE("/products/:product_id", mc.DeleteProduct)
		merchants.GET("/orders", mc.ShowAllOrders)
		merchants.PUT("/orders/:order_id", mc.ProcessOrder)
	}

}
