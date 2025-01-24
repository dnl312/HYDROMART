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

// @Summary     Create new order
// @Description Create new order
// @Tags        Order
// @Accept      json
// @Produce     json
// @Param       request body model.CreateOrder true "User Order Create"
// @Success     201 {object} map[string]string
// @failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /orders/create [post]
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
	return ctx.JSON(http.StatusCreated, map[string]string{
		"message": "Order created success",
	})
}

// @Summary     Delete order
// @Description Delete order
// @Tags        Order
// @Accept      json
// @Produce     json
// @Param       request body model.DeleteOrder true "User Order Delete"
// @Success     200 {object} map[string]string
// @failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /orders/order-delete [delete]
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

// @Summary     Get all orders
// @Description Get all orders
// @Tags        Order
// @Accept      json
// @Produce     json
// @Param       status query string true "Order status"
// @Success     200 {object} map[string]string
// @failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /orders/order-list [get]
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

// @Summary     Update order
// @Description Update order
// @Tags        Order
// @Accept      json
// @Produce     json
// @Param       request body model.UpdateOrder true "User Order Update"
// @Success     200 {object} map[string]string
// @failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /orders/order-update [put]
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

// @Summary     Top up user deposit
// @Description Top up user deposit
// @Tags        Order
// @Accept      json
// @Produce     json
// @Param       request body model.TopUp true "User Top Up"
// @Success     200 {object} map[string]string
// @failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /orders/topup [post]
func (u OrderController) TopUp(ctx echo.Context) error {
	var req model.TopUp
	if err := ctx.Bind(&req); err != nil {
		log.Print(err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
	}
	token := ctx.Request().Header.Get("Authorization")
	md := metadata.Pairs("Authorization", token)
	ctxWithToken := metadata.NewOutgoingContext(context.Background(), md)

	serviceCtx, cancel := context.WithTimeout(ctxWithToken, 10*time.Second)
	defer cancel()

	r, err := u.Client.CreateTopUp(serviceCtx, &pb.TopUpUserDepositRequest{Amount: float32(req.Amount)})
	if err != nil {
		log.Printf("could not top up: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "top up error"})
	}
	log.Printf("top up Response: %v", r)

	return ctx.JSON(http.StatusOK, r)
}

// @Summary     Update deposit cron
// @Description Update deposit cron
// @Tags        Order
// @Accept      json
// @Produce     json
// @Success     200 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /orders/update-deposit [get]
func (u OrderController) UpdateDepositCron(ctx echo.Context) error {
	serviceCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := u.Client.SchedulerUpdateDeposit(serviceCtx, &pb.SchedulerUpdateDepositRequest{})
	if err != nil {
		log.Printf("could not update deposit: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "update deposit error"})
	}
	log.Printf("update deposit Response: %v", r)

	return ctx.JSON(http.StatusOK, r)
}