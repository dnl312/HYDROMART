package repo

import (
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