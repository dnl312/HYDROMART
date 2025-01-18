package model

type Product struct {
	ProductID   string `gorm:"primaryKey"`
	MerchantID  string `gorm:"unique"`
	ProductName string `gorm:"not null"`
	Price       string `gorm:"not null"`
	Stock       int    `gorm:"not null"`
	Category    string `gorm:"not null"`
	Merchant    User   `gorm:"foreignKey:UserID"`
}
