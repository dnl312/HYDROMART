package model

type Transaction struct {
	TransactionID string    `gorm:"unique"`
	UserID        string    `gorm:"not null"`
	ProductID     string    `gorm:"not null"`
	Qty           int       `gorm:"not null"`
	Amount        float64   `gorm:"not null"`
	Status        string    `gorm:"not null"`
}
