package controller

import (
	"context"
	"testing"
	"time"
	"user/model"
	"user/pb"
	"user/repo"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	userRepository = repo.NewMockOrderRepository()
	userController = NewOrderController(&userRepository)
)

var (
	modelOrder = model.TransactionLs{
		TransactionID:  primitive.NewObjectID().Hex(),
		UserID: "123",
		ProductID: "456",
		Qty: 1,
		Amount: 1000,
		Status: "ORDER CREATED",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestGetAllOrdersWithStatus(t *testing.T) {
	pbRequest := &pb.GetAllOrdersWithStatusRequest{
		Status: "ORDER CREATED",
	}

	order := []model.TransactionLs{
		modelOrder,
	}

	userID := "123"

	userRepository.Mock.On("GetOrder", userID, pbRequest.Status).Return(order, nil)


	pbResponse, err := userController.GetAllOrdersWithStatus(context.Background(), pbRequest)
	
    assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
}

func TestCreateOrder(t *testing.T) {
	pbRequest := &pb.CreateOrderRequest{
		ProductId: "456",
		Qty: 1,
	}

	userRepository.Mock.On("CreateOrder", pbRequest).Return(nil)

	pbResponse, err := userController.CreateOrder(context.Background(), pbRequest)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
}

func TestDeleteOrder(t *testing.T) {
	pbRequest := &pb.DeleteOrderRequest{
		TransactionId: "123-INV",
	}

	userRepository.Mock.On("DeleteOrder", pbRequest.TransactionId).Return(nil)

	pbResponse, err := userController.DeleteOrder(context.Background(), pbRequest)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
}

func TestUpdateOrder(t *testing.T) {
	pbRequest := &pb.UpdateOrderRequest{
		TransactionId: "123-INV",
		Qty: 1,
	}

	userRepository.Mock.On("UpdateOrder", pbRequest).Return(nil)

	pbResponse, err := userController.UpdateOrder(context.Background(), pbRequest)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
}
