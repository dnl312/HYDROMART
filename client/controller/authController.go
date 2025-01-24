package controller

import (
	"client/model"
	pb "client/pb/authpb"
	"context"
	"time"

	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	Client pb.AuthServiceClient
}

func NewAuthController(client pb.AuthServiceClient) AuthController {
	return AuthController{
		Client: client,
	}
}

// @Summary     Login a user
// @Description Login a user
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       request body model.LoginRequest true "User login details"
// @Success     200 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /users/login [post]
func (u AuthController) LoginUser (ctx echo.Context) error{
		var req model.LoginRequest
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
		}
		serviceCtx, cancel:= context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		 log.Printf("could not login: %v", req)

		r, err := u.Client.LoginUser(serviceCtx, &pb.LoginRequest{Username: req.Username, Password: req.Password})
		if err != nil {
			log.Printf("could not login2: %v", err)
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "login error"})
		}
		log.Printf("Login Response: %s", r.GetToken())

		return ctx.JSON(http.StatusOK, map[string]string{
			"token": r.Token,
		})
}

// @Summary     Register a new user
// @Description Register a new user with the role 'USER'
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       request body model.RegisterUser true "User registration details"
// @Success     201 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /users/register [post]
func (u AuthController) RegisterUser (ctx echo.Context) error{
	var req model.RegisterUser
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
	}

	serviceCtx, cancel:= context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := u.Client.RegisterUser(serviceCtx, &pb.RegisterRequest{
		Username: req.Username, 
		Password: req.Password, 
		Email: req.Email,
		Address: req.Address,
		Role: "USER",
	})
	if err != nil {
		log.Printf("could not register: %v", err)
		log.Printf("could not register: %v", req)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "registration failed"})
	}
	log.Printf("Register Response: %s", r.GetMessage())

	return ctx.JSON(http.StatusCreated, map[string]string{
		"message": r.Message,
	})
}

// @Summary     Register a new user
// @Description Register a new user with the role 'MERCHANT'
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       request body model.RegisterUser true "User registration details"
// @Success     201 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /users/register/merchant [post]
func (u AuthController) RegisterMerchant (ctx echo.Context) error{
	var req model.RegisterUser
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
	}

	serviceCtx, cancel:= context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := u.Client.RegisterUser(serviceCtx, &pb.RegisterRequest{
		Username: req.Username, 
		Password: req.Password, 
		Email: req.Email,
		Address: req.Address,
		Role: "MERCHANT",
	})
	if err != nil {
		log.Printf("could not register: %v", err)
		log.Printf("could not register: %v", req)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "registration failed"})
	}
	log.Printf("Register Response: %s", r.GetMessage())

	return ctx.JSON(http.StatusCreated, map[string]string{
		"message": r.Message,
	})
}