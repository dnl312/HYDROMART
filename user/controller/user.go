package controller

import (
	"context"
	"log"
	"user/middleware"
	"user/model"
	"user/pb"
	"user/repo"
	"user/utils"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (u *User) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	token, err := middleware.GetTokenStringFromContext(ctx)
	if err != nil {
		log.Println("Received nil request, returning error")
		return nil, err
	}

	userClaims, err := utils.RecoverUser(token)
	if err != nil {
		return nil, err
	}

	product, err := u.Repository.GetProductByID(req.ProductId)
	if err != nil {
		return nil, err
	}

	totalPrice := product.Price * float64(req.Qty)

	user, err := u.Repository.ValidateUser(userClaims.UserID)
	if err != nil {
		return nil, err
	}

	if user.Deposit < totalPrice {
		return nil, status.Errorf(codes.Internal, "user's deposit is low, please top up: %v", err) 
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
		return nil, status.Errorf(codes.Internal, "Create order failed: %v", err) 
	}

	err = u.Repository.UpdateDeposit(order)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "update user's deposit failed: %v", err) 
	}

	return &pb.CreateOrderResponse{
		Message: "Order created successfully",
	}, nil
}

func (u *User) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	token, err := utils.GetTokenStringFromContext(ctx)
	if err != nil {
		return nil, err
	}

	_, err = utils.RecoverUser(token)
	if err != nil {
		return nil, err
	}

	orderID := req.TransactionId

	if orderID == "" {
		return nil, err
	}

	err = u.Repository.DeleteOrder(orderID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return &pb.DeleteOrderResponse{
		Message: "Order deleted successfully",
	}, nil
}
