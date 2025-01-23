package model

import "time"

type Transaction struct {
	TransactionID string  `gorm:"unique"`
	UserID        string  `gorm:"not null"`
	ProductID     string  `gorm:"not null"`
	Qty           int     `gorm:"not null"`
	Amount        float64 `gorm:"not null"`
	Status        string  `gorm:"not null"`
}

type TransactionLs struct {
	TransactionID string    `gorm:"unique"`
	UserID        string    `gorm:"not null"`
	ProductID     string    `gorm:"not null"`
	Qty           int       `gorm:"not null"`
	Amount        float64   `gorm:"not null"`
	Status        string    `gorm:"not null"`
	CreatedAt     time.Time `gorm:"not null"`
	UpdatedAt     time.Time `gorm:"not null"`
}
