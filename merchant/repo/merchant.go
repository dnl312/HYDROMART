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

func (u *MerchantRepository) GetAllProduct(merchantID string) (*[]model.Product, error) {
	var products []model.Product
	result := u.DB.Where("merchant_id ?", merchantID).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return &products, nil
}

func (u *MerchantRepository) AddProduct(productPtr *model.Product) error {
	result := u.DB.Create(productPtr)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *MerchantRepository) UpdateProduct(productPtr *model.Product) error {
	result := u.DB.Save(productPtr)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *MerchantRepository) DeleteProduct(productID string) error {
	result := u.DB.Delete(&model.Product{}, productID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
