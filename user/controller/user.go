package controller

import (
	//pb "user/pb"
	"context"
	"user/model"
	pb "user/pb"
	"user/repo"
	"user/utils"
)

type User struct {
	pb.UnimplementedOrderServer
	Repository repo.UserInterface
	Client     pb.GetUserServiceClient
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

	userData, err := u.Client.GetUserByID(ctx, &pb.GetUserByIDRequest{UserId: req.UserId})
	if err != nil {
		return nil, err
	}

	if userData.Deposit < totalPrice {
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

func (u *User) GetAllOrders(ctx context.Context, req *pb.GetAllOrdersRequest) (*pb.GetAllOrdersResponse, error) {
	tokenString, err := utils.GetTokenStringFromContext(ctx)
	if err != nil {
		return nil, err
	}
	user, err := utils.RecoverUser(tokenString)
	if err != nil {
		return nil, err
	}

	orderID := user["user_id"].(string)

	allOrder, err := u.Repository.GetOrder(orderID)
	if err != nil {
		return nil, err
	}

	allOrderResponse := pb.GetAllOrdersResponse{
		Orders: []*pb.OrderResponse{},
	}

	for _, order := range *allOrder {
		orderResponse := pb.OrderResponse{
			TransactionId: order.TransactionID,
			UserId:        order.UserID,
			ProductId:     order.ProductID,
			Qty:           int32(order.Qty),
			Amount:        float32(order.Amount),
			Status:        order.Status,
		}
		allOrderResponse.Orders = append(allOrderResponse.Orders, &orderResponse)
	}
	return &allOrderResponse, nil
}

func (u *User) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	tokenString, err := utils.GetTokenStringFromContext(ctx)
	if err != nil {
		return nil, err
	}
	user, err := utils.RecoverUser(tokenString)
	if err != nil {
		return nil, err
	}

	orderID := user["user_id"].(string)

	order := model.Transaction{
		TransactionID: orderID,
		ProductID:     req.ProductId,
		Qty:           int(req.Qty),
	}

	err = u.Repository.UpdateOrder(&order)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateOrderResponse{
		Message: "Order has been updated",
	}, nil

}

func (u *User) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	err := u.Repository.DeleteOrder(req.TransactionId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteOrderResponse{
		Message: "Order delete succesfully",
	}, nil
}
