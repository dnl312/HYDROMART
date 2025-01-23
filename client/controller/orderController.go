package controller

import (
	"client/model"
	pb "client/pb/userpb"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/metadata"
)

type OrderController struct {
	Client pb.OrderClient
}

func NewOrderController(client pb.OrderClient) OrderController {
	return OrderController{
		Client: client,
	}
}

func (u OrderController) CreateOrder(ctx echo.Context) error {
	var req model.CreateOrder
	if err := ctx.Bind(&req); err != nil {
		log.Print(err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
	}
	token := ctx.Request().Header.Get("Authorization")
	md := metadata.Pairs("Authorization", token)
	ctxWithToken := metadata.NewOutgoingContext(context.Background(), md)

	serviceCtx, cancel := context.WithTimeout(ctxWithToken, 10*time.Second)
	defer cancel()

	log.Printf("could not login: %v", req)

	_, err := u.Client.CreateOrder(serviceCtx, &pb.CreateOrderRequest{ProductId: req.ProductID, Qty: int32(req.Qty)})
	if err != nil {
		log.Printf("could not create order: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "create order error"})
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Order created success",
	})
}

func (u OrderController) DeleteOrder(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")
	md := metadata.Pairs("Authorization", token)
	ctxWithToken := metadata.NewOutgoingContext(context.Background(), md)

	serviceCtx, cancel := context.WithTimeout(ctxWithToken, 10*time.Second)
	defer cancel()

	var req model.DeleteOrder
	if err := ctx.Bind(&req); err != nil {
		log.Print(err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
	}

	r, err := u.Client.DeleteOrder(serviceCtx, &pb.DeleteOrderRequest{
		TransactionId: req.TransactionID,
	})

	if err != nil {
		log.Printf("could not delete transaction: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "delete transaction error"})
	}
	log.Printf("update product delete: %v", r)

	return ctx.JSON(http.StatusOK, r)

}

func (u OrderController) GetAllOrders(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")
	md := metadata.Pairs("Authorization", token)
	ctxWithToken := metadata.NewOutgoingContext(context.Background(), md)

	serviceCtx, cancel := context.WithTimeout(ctxWithToken, 10*time.Second)
	defer cancel()

	var req pb.GetAllOrdersWithStatusRequest
	if err := ctx.Bind(&req); err != nil {
		log.Print(err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
	}

	r, err := u.Client.GetAllOrdersWithStatus(serviceCtx, &pb.GetAllOrdersWithStatusRequest{Status: req.Status})
	if err != nil {
		log.Printf("could not show all order: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "show all order error"})
	}
	log.Printf("show all order Response: %v", r)

	return ctx.JSON(http.StatusOK, r)
}


func (u OrderController) UpdateOrder(ctx echo.Context) error {
	var req model.UpdateOrder
	if err := ctx.Bind(&req); err != nil {
		log.Print(err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
	}
	token := ctx.Request().Header.Get("Authorization")
	md := metadata.Pairs("Authorization", token)
	ctxWithToken := metadata.NewOutgoingContext(context.Background(), md)

	serviceCtx, cancel := context.WithTimeout(ctxWithToken, 10*time.Second)
	defer cancel()

	r, err := u.Client.UpdateOrder(serviceCtx, &pb.UpdateOrderRequest{TransactionId: req.TransactionID, Qty: int32(req.Qty)})
	if err != nil {
		log.Printf("could not update order: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "update order error"})
	}
	log.Printf("update order Response: %v", r)

	return ctx.JSON(http.StatusOK, r)
}

func (u OrderController) UpdateDepositCron(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")
	md := metadata.Pairs("Authorization", token)
	ctxWithToken := metadata.NewOutgoingContext(context.Background(), md)

	serviceCtx, cancel := context.WithTimeout(ctxWithToken, 10*time.Second)
	defer cancel()

	r, err := u.Client.SchedulerUpdateDeposit(serviceCtx, &pb.SchedulerUpdateDepositRequest{})
	if err != nil {
		log.Printf("could not update deposit: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "update deposit error"})
	}
	log.Printf("update deposit Response: %v", r)

	return ctx.JSON(http.StatusOK, r)
}