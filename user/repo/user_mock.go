package repo

import (
	"user/model"

	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func NewMockOrderRepository() MockOrderRepository {
	return MockOrderRepository{}
}

func (m *MockOrderRepository) ValidateUser(uuid string) (model.User, error){
	args := m.Called(uuid)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockOrderRepository) UpdateDeposit(order model.Transaction) error{
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderRepository) GetOrder(userID string, status string) (*[]model.TransactionLs, error){
	args := m.Called()
	return args.Get(0).(*[]model.TransactionLs), args.Error(1)
}

func (m *MockOrderRepository) GetOrderByID(orderID string) (model.Transaction, error){
	args := m.Called(orderID)
	return args.Get(0).(model.Transaction), args.Error(1)
}

func (m *MockOrderRepository) GetProductByID(productID string) (model.Product, error){
	args := m.Called(productID)
	return args.Get(0).(model.Product), args.Error(1)
}

func (m *MockOrderRepository) CreateOrder(order model.Transaction) error{
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderRepository) DeleteOrder(orderID string) error{
	args := m.Called(orderID)
	return args.Error(0)
}
func (m *MockOrderRepository) UpdateOrder(order model.Transaction) error{
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderRepository) InsertIntoTopUpTemp(topup_id string, user_id string) error{
	args := m.Called(topup_id, user_id)
	return args.Error(0)
}