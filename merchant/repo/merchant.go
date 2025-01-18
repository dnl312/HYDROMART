package repo

import (
	"merchant/model"

	"gorm.io/gorm"
)

type MerchantRepository struct {
	DB *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) MerchantRepository {
	return MerchantRepository{
		DB: db,
	}
}

func (u *MerchantRepository) GetAllProduct() (*[]model.Product, error) {
	var products []model.Product
	result := u.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return &products, nil
}
