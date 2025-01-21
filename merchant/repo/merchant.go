package repo

import (
	"errors"
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

	result := u.DB.Table("products_hydromart").Where("merchant_id = ?", merchantID).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return &products, nil
}

func (u *MerchantRepository) AddProduct(productPtr *model.Product) error {
	result := u.DB.Table("products_hydromart").Create(productPtr)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *MerchantRepository) UpdateProduct(productPtr *model.Product) error {
	result := u.DB.Table("products_hydromart").Where("merchant_id = ?", productPtr.MerchantID).Save(productPtr)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *MerchantRepository) DeleteProduct(productID, merchantID string) error {
	result := u.DB.Table("products_hydromart").Where("product_id = ?", productID).Where("merchant_id = ?", merchantID).Delete(&model.Product{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}
