package model

type Product struct {
	ProductID 	string `gorm:"unique"`
	MerchantID 	string `gorm:"not null"`
	ProductName string `gorm:"not null"`
	Price float64 `gorm:"not null"`
	Stock int `gorm:"not null"`
	Category string `gorm:"not null"`
}