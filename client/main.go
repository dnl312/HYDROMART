package main

import (
	"client/config"
	"client/controller"
	helpers "client/middleware"
	"client/router"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Validator = &helpers.CustomValidator{NewValidator: validator.New()}
	e.Use(middleware.Logger(), middleware.Recover(), middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	authClientConn, authClient := config.InitAuthServiceClient()
	defer authClientConn.Close()

	merchantClientConn, merchantClient := config.InitMerchantServiceClient()
	defer merchantClientConn.Close()
  
  orderClientConn, orderClient := config.InitOrderServiceClient()
	defer orderClientConn.Close()

  orderController := controller.NewOrderController(orderClient)
  merchantController := controller.NewMerchantController(merchantClient)
	authController := controller.NewAuthController(authClient)
  
  router.Echo(e, authController, merchantController, orderController)

	e.Logger.Fatal(e.Start(":8080"))
}
