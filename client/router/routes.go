package router

import (
	"client/controller"

	"github.com/labstack/echo/v4"
)

func Echo(e *echo.Echo, uc controller.AuthController) {
	users := e.Group("/users")
	{		
		users.POST("/login", uc.LoginUser)
		users.POST("/register", uc.RegisterUser)
	}

}