package controller

import (
	//pb "user/pb"
	"context"
	"user/model"
	"user/pb"
	"user/repo"
	"user/utils"

	"gorm.io/gorm"
)

type User struct {
	//pb.UnimplementedMerchantServiceServer
	Repository repo.UserInterface
}

func NewMerchantController(r repo.UserInterface) User {
	return User{
		Repository: r,
	}
}

func (u *User) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	token, err := utils.GetTokenStringFromContext(ctx)
	if err != nil {
		return nil, err
	}

	userClaims, err := utils.RecoverUser(token)
	if err != nil {
		return nil, err
	}

	user_ID := userClaims["user_id"].(string)

	if req.Qty <= 0 {
		return nil, err
	}

	product, err := u.Repository.GetProductByID(req.ProductId)
	if err != nil {
		return nil, err
	}

	totalPrice := product.Price * float64(req.Qty)

	userAmount := userClaims["amount"].(float64)
	if userAmount < totalPrice {
		return nil, err
	}

	order := &model.Transaction{
		UserID:    user_ID,
		ProductID: req.ProductId,
		Qty:       int(req.Qty),
		Amount:    totalPrice,
		Status:    "order created",
	}

	err = u.Repository.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	return &pb.CreateOrderResponse{
		Message: "Order created successfully",
	}, nil
}

func (u *User) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderRequest, error) {
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
	return &pb.CreateOrderResponse{
		Message: "Transaction deleted successfully",
	}, nil
}
