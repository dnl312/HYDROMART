package controller

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"
	"user/middleware"
	"user/model"
	"user/pb"
	"user/repo"
	"user/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	pb.UnimplementedOrderServer
	Repository repo.UserInterface
}

func NewOrderController(r repo.UserInterface) User {
	return User{
		Repository: r,
	}
}

func (u *User) GetAllOrdersWithStatus(ctx context.Context, req *pb.GetAllOrdersWithStatusRequest) (*pb.GetAllOrdersWithStatusResponse, error) {
	token, err := middleware.GetTokenStringFromContext(ctx)
	if err != nil {
		return &pb.GetAllOrdersWithStatusResponse{
			Orders: []*pb.OrderResponse{},
		}, nil
	}

	user, err := utils.RecoverUser(token)
	if err != nil {
		return &pb.GetAllOrdersWithStatusResponse{
			Orders: []*pb.OrderResponse{},
		}, nil
	}

	orders, err := u.Repository.GetOrder(user.UserID, req.Status)
	if err != nil {
		return &pb.GetAllOrdersWithStatusResponse{
			Orders: []*pb.OrderResponse{},
		}, nil
	}

	orderResponse := pb.GetAllOrdersWithStatusResponse{
		Orders: []*pb.OrderResponse{},
	}
	for _, order := range *orders {
		orderResponse.Orders = append(orderResponse.Orders, &pb.OrderResponse{
			TransactionId: order.TransactionID,
			UserId:        order.UserID,
			ProductId:     order.ProductID,
			Qty:           int32(order.Qty),
			Amount:        float32(order.Amount),
			Status:        order.Status,
			CreatedAt:    order.CreatedAt.String(),
			UpdatedAt:   order.UpdatedAt.String(),
		})
	}
	return &orderResponse, nil
}

func (u *User) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	token, err := middleware.GetTokenStringFromContext(ctx)
	if err != nil {
		return &pb.CreateOrderResponse{Message: "not authorized",}, nil
	}

	userClaims, err := utils.RecoverUser(token)
	if err != nil {
		return &pb.CreateOrderResponse{Message: "JWT failed",}, nil
	}

	product, err := u.Repository.GetProductByID(req.ProductId)
	if err != nil || err == gorm.ErrRecordNotFound {
		return &pb.CreateOrderResponse{Message: "product not found",}, nil
	}

	totalPrice := product.Price * float64(req.Qty)

	user, err := u.Repository.ValidateUser(userClaims.UserID)
	if err != nil || err == gorm.ErrRecordNotFound {
		return &pb.CreateOrderResponse{Message: "user not found",}, nil
	}

	if user.Deposit < totalPrice {
		return &pb.CreateOrderResponse{Message: "user's deposit is low, please top up",},nil
	}

	order := model.Transaction{
		TransactionID: uuid.New().String(),
		UserID:    userClaims.UserID,
		ProductID: req.ProductId,
		Qty:       int(req.Qty),
		Amount:    totalPrice,
		Status:    "ORDER CREATED",

	}
	log.Printf("Received order: %v", order)

	err = u.Repository.CreateOrder(order)
	if err != nil {
		return &pb.CreateOrderResponse{Message: "Create order failed",},nil
	}

	err = u.Repository.UpdateDeposit(order)
	if err != nil {
		return &pb.CreateOrderResponse{Message: "update user's deposit failed",}, nil
	}

	return &pb.CreateOrderResponse{
		Message: "Order created successfully",
	}, nil
}

func (u *User) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	token, err := utils.GetTokenStringFromContext(ctx)
	if err != nil {
		return &pb.DeleteOrderResponse{Message: "not authorized",}, nil
	}

	_, err = utils.RecoverUser(token)
	if err != nil {
		return &pb.DeleteOrderResponse{Message: "JWT failed",}, nil
	}

	orderID := req.TransactionId

	if orderID == "" {
		return &pb.DeleteOrderResponse{Message: "transaction id not found",}, nil
	}

	err = u.Repository.DeleteOrder(orderID)
	if err != nil || err == gorm.ErrRecordNotFound {
		return &pb.DeleteOrderResponse{Message: "transaction not found",}, nil
	}
	return &pb.DeleteOrderResponse{Message: "Order deleted successfully",}, nil
}

func (u *User) UpdateOrder (ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	token, err := utils.GetTokenStringFromContext(ctx)
	if err != nil {
		return &pb.UpdateOrderResponse{Message: "not authorized",}, nil
	}

	user, err := utils.RecoverUser(token)
	if err != nil {
		return &pb.UpdateOrderResponse{Message: "JWT failed",}, nil
	}

	trx, err := u.Repository.GetOrderByID(req.TransactionId)
	if err != nil || err == gorm.ErrRecordNotFound {
		return &pb.UpdateOrderResponse{Message: "transaction not found",}, nil
	}

	product, err := u.Repository.GetProductByID(trx.ProductID)
	if err != nil || err == gorm.ErrRecordNotFound {
		return &pb.UpdateOrderResponse{Message: "product not found",}, nil
	}

	totalPrice := product.Price * float64(req.Qty)

	userClaims, err := u.Repository.ValidateUser(user.UserID)
	if err != nil || err == gorm.ErrRecordNotFound {
		return &pb.UpdateOrderResponse{Message: "user not found",}, nil
	}

	if userClaims.Deposit < totalPrice {
		return &pb.UpdateOrderResponse{
			Message: "user's deposit is low, please top up",
		}, nil
	}

	order := model.Transaction{
		TransactionID: req.TransactionId,
		UserID: user.UserID,
		Qty: int(req.Qty),
		Amount: totalPrice,
	}

	err = u.Repository.UpdateOrder(order)
	if err != nil {
		return nil, err
	}

	if int(req.Qty) > trx.Qty || int(req.Qty) < trx.Qty{
		totalSelisih := int(req.Qty) - trx.Qty
		order.Amount = product.Price * float64(totalSelisih)

		err = u.Repository.UpdateDeposit(order)
		if err != nil {
			return &pb.UpdateOrderResponse{Message: "update user's deposit failed",}, nil
		}
	}

	return &pb.UpdateOrderResponse{Message: "Order updated successfully",}, nil
}

func (u *User) CreateTopUp(ctx context.Context, req *pb.TopUpUserDepositRequest) (*pb.TopUpUserDepositResponse, error) {
	token, err := utils.GetTokenStringFromContext(ctx)
	if err != nil {
		return &pb.TopUpUserDepositResponse{}, err
	}

	user, err := utils.RecoverUser(token)
	if err != nil {
		return &pb.TopUpUserDepositResponse{}, err
	}
	newUrl := os.Getenv("MIDTRANS_URL") + "/snap/v1/transactions"
	headers := map[string]string{
		"authorization":  os.Getenv("MIDTRANS_APIKEY"),
		"content-type": "application/json",
	}

	currentTime := time.Now()
	timestampString := currentTime.Format("20060102150405")

	requestBody := map[string]interface{}{
		"transaction_details": map[string]interface{}{
			"order_id":    "topup_" + timestampString,
			"gross_amount": req.Amount,
		},
		"credit_card": map[string]interface{}{
			"secure": false,
		},
	}
	
	jsonBody, err := json.Marshal(requestBody)
    if err != nil {
        return &pb.TopUpUserDepositResponse{}, err
    }

	var midtransResponse pb.TopUpUserDepositResponse
	payload := strings.NewReader(string(jsonBody))
	response, err := utils.RequestPOST(newUrl, headers, payload)
	if err != nil {
		return &pb.TopUpUserDepositResponse{}, err
	}

	if len(response) == 0 {
		return &pb.TopUpUserDepositResponse{}, err
	}

	err = json.Unmarshal(response, &midtransResponse)
	if err != nil {
		return &pb.TopUpUserDepositResponse{}, err
	}

	err = u.Repository.InsertIntoTopUpTemp("topup_" + timestampString, user.UserID)
	if err != nil {
		return &pb.TopUpUserDepositResponse{}, err
	}

	return &pb.TopUpUserDepositResponse{
		Token:        midtransResponse.Token,
		RedirectUrl: midtransResponse.RedirectUrl,
	}, nil
}