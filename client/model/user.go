package model

type User struct {
	UserID 	string `gorm:"unique"`
	Username 	string `gorm:"unique"`
	Email string `gorm:"not null"`
	Password string `gorm:"not null"`
	Address string `gorm:"not null"`
	Role string `gorm:"not null"`
	Deposit float64 `gorm:"not null"`
}