package controller

import (
	"client/model"
	pb "client/pb/merchantpb"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/metadata"
)

type MerchantController struct {
	Client pb.MerchantServiceClient
}

func NewMerchantController(client pb.MerchantServiceClient) MerchantController {
	return MerchantController{
		Client: client,
	}
}

func (mc MerchantController) ShowAllProducts(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")
	md := metadata.Pairs("Authorization", token)
	ctxWithToken := metadata.NewOutgoingContext(context.Background(), md)

	serviceCtx, cancel := context.WithTimeout(ctxWithToken, 10*time.Second)
	defer cancel()

	r, err := mc.Client.ShowAllProducts(serviceCtx, &pb.ShowAllProductRequest{})
	if err != nil {
		log.Printf("could not show all product: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "show all product error"})
	}
	log.Printf("show all product Response: %v", r)

	return ctx.JSON(http.StatusOK, r)
}

func (mc MerchantController) AddProduct(ctx echo.Context) error {
	var req model.AddProductRequest
	if ctx.Bind(&req) != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid JSON request")
	}

	token := ctx.Request().Header.Get("Authorization")
	md := metadata.Pairs("Authorization", token)
	ctxWithToken := metadata.NewOutgoingContext(context.Background(), md)

	serviceCtx, cancel := context.WithTimeout(ctxWithToken, 10*time.Second)
	defer cancel()

	product := pb.Product{
		Name:     req.Name,
		Price:    req.Price,
		Stock:    int32(req.Stock),
		Category: req.Category,
	}

	r, err := mc.Client.AddProduct(serviceCtx, &pb.AddProductRequest{
		Product: &product,
	})
	if err != nil {
		log.Printf("could not add product: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "add product error"})
	}
	log.Printf("add product Response: %v", r)

	return ctx.JSON(http.StatusOK, r)
}

func (mc MerchantController) UpdateProduct(ctx echo.Context) error {
	var req model.Product
	if ctx.Bind(&req) != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid JSON request")
	}

	token := ctx.Request().Header.Get("Authorization")
	md := metadata.Pairs("Authorization", token)
	ctxWithToken := metadata.NewOutgoingContext(context.Background(), md)

	serviceCtx, cancel := context.WithTimeout(ctxWithToken, 10*time.Second)
	defer cancel()

	product := pb.Product{
		Id:       ctx.Param("product_id"),
		Name:     req.ProductName,
		Price:    req.Price,
		Stock:    int32(req.Stock),
		Category: req.Category,
	}

	r, err := mc.Client.UpdateProduct(serviceCtx, &pb.UpdateProductRequest{
		Product: &product,
	})
	if err != nil {
		log.Printf("could not update product: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "update product error"})
	}
	log.Printf("update product Response: %v", r)

	return ctx.JSON(http.StatusOK, r)
}

func (mc MerchantController) DeleteProduct(ctx echo.Context) error {

	token := ctx.Request().Header.Get("Authorization")
	md := metadata.Pairs("Authorization", token)
	ctxWithToken := metadata.NewOutgoingContext(context.Background(), md)

	serviceCtx, cancel := context.WithTimeout(ctxWithToken, 10*time.Second)
	defer cancel()

	r, err := mc.Client.DeleteProduct(serviceCtx, &pb.DeleteProductRequest{
		ProductId: ctx.Param("product_id"),
	})
	if err != nil {
		log.Printf("could not delete product: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "delete product error"})
	}
	log.Printf("update product delete: %v", r)

	return ctx.JSON(http.StatusOK, r)
}

func (mc MerchantController) ShowAllOrders(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")
	md := metadata.Pairs("Authorization", token)
	ctxWithToken := metadata.NewOutgoingContext(context.Background(), md)

	serviceCtx, cancel := context.WithTimeout(ctxWithToken, 10*time.Second)
	defer cancel()

	r, err := mc.Client.ShowAllOrders(serviceCtx, &pb.ShowAllOrderRequest{})
	if err != nil {
		log.Printf("could not show all orders: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "show all orders error"})
	}
	log.Printf("show all order Response: %v", r)

	return ctx.JSON(http.StatusOK, r)
}

func (mc MerchantController) ProcessOrder(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")
	md := metadata.Pairs("Authorization", token)
	ctxWithToken := metadata.NewOutgoingContext(context.Background(), md)

	serviceCtx, cancel := context.WithTimeout(ctxWithToken, 10*time.Second)
	defer cancel()

	r, err := mc.Client.ProcessOrder(serviceCtx, &pb.ProcessOrderRequest{
		OrderId: ctx.Param("order_id"),
	})
	if err != nil {
		log.Printf("could not update order: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "couldn't process order"})
	}
	log.Printf("update process order response: %v", r)

	return ctx.JSON(http.StatusOK, r)
}
