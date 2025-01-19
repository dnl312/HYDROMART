package controller

import (
	"client/model"
	pb "client/pb/merchantpb"
	"context"
	"time"

	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type MerchantController struct {
	Client pb.MerchantServiceClient
}

func NewMerchantController(client pb.MerchantServiceClient) MerchantController {
	return MerchantController{
		Client: client,
	}
}

// custom jwt claim
type jwtCustomClaims struct {
	ID    string `json:"user_id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func (u MerchantController) ShowAllProduct(ctx echo.Context) error {
	token := ctx.Get("Authorization").(*jwt.Token)
	claims := token.Claims.(*jwtCustomClaims)

	if claims.Role != "merchant" {
		return ctx.JSON(http.StatusUnauthorized, "unauthorized")
	}

	serviceCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := u.Client.ShowAllProduct(serviceCtx, &pb.ShowAllProductRequest{MerchantId: claims.ID})
	if err != nil {
		log.Printf("could not show all product: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "show all product error"})
	}
	log.Printf("show all product Response: %v", r)

	return ctx.JSON(http.StatusOK, r)
}

func (u MerchantController) AddProduct(ctx echo.Context) error {
	var req model.Product
	if ctx.Bind(&req) != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid JSON request")
	}

	token := ctx.Get("Authorization").(*jwt.Token)
	claims := token.Claims.(*jwtCustomClaims)

	if claims.Role != "merchant" {
		return ctx.JSON(http.StatusUnauthorized, "unauthorized")
	}

	serviceCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	product := pb.Product{
		Id:         req.ProductID,
		MerchantId: req.MerchantID,
		Name:       req.ProductName,
		Price:      req.Price,
		Stock:      int32(req.Stock),
		Category:   req.Category,
	}

	r, err := u.Client.AddProduct(serviceCtx, &pb.AddProductRequest{
		MerchantId: claims.ID,
		Product:    &product,
	})
	if err != nil {
		log.Printf("could not add product: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "add product error"})
	}
	log.Printf("add product Response: %v", r)

	return ctx.JSON(http.StatusOK, r)
}

func (u MerchantController) UpdateProduct(ctx echo.Context) error {
	var req model.Product
	if ctx.Bind(&req) != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid JSON request")
	}

	token := ctx.Get("Authorization").(*jwt.Token)
	claims := token.Claims.(*jwtCustomClaims)

	if claims.Role != "merchant" {
		return ctx.JSON(http.StatusUnauthorized, "unauthorized")
	}

	serviceCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	product := pb.Product{
		Id:         req.ProductID,
		MerchantId: req.MerchantID,
		Name:       req.ProductName,
		Price:      req.Price,
		Stock:      int32(req.Stock),
		Category:   req.Category,
	}

	r, err := u.Client.UpdateProduct(serviceCtx, &pb.UpdateProductRequest{
		MerchantId: claims.ID,
		Product:    &product,
	})
	if err != nil {
		log.Printf("could not update product: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "update product error"})
	}
	log.Printf("update product Response: %v", r)

	return ctx.JSON(http.StatusOK, r)
}

func (u MerchantController) DeleteProduct(ctx echo.Context) error {
	var req model.Product
	if ctx.Bind(&req) != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid JSON request")
	}

	token := ctx.Get("Authorization").(*jwt.Token)
	claims := token.Claims.(*jwtCustomClaims)

	if claims.Role != "merchant" {
		return ctx.JSON(http.StatusUnauthorized, "unauthorized")
	}

	serviceCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := u.Client.DeleteProduct(serviceCtx, &pb.DeleteProductRequest{
		MerchantId: req.MerchantID,
		ProductId:  req.ProductID,
	})
	if err != nil {
		log.Printf("could not update product: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "update product error"})
	}
	log.Printf("update product Response: %v", r)

	return ctx.JSON(http.StatusOK, r)
}
