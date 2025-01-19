package controller

import (
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
