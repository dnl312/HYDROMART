package repo

import (
	"user/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		DB: db,
	}
}

func (u *UserRepository) GetProductByID(productID string) (*model.Product, error) {
	var product model.Product
	result := u.DB.Where("id = ?", productID).First(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (u *UserRepository) CreateOrder(order *model.Transaction) error {
	result := u.DB.Create(order)
	if result != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepository) GetOrder(userID string) (*[]model.Transaction, error) {
	var Order []model.Transaction
	result := u.DB.Where("user_id = ?", userID).Find(&Order)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Order, nil
}

func (u *UserRepository) UpdateOrder(order *model.Transaction) error {
	result := u.DB.Save(order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepository) DeleteOrder(orderID string) error {
	var Order model.Transaction
	result := u.DB.Where("transaction_id = ?", orderID).Delete(&Order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
