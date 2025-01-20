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
