package repo

import (
	"merchant/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		DB: db,
	}
}

func (u *UserRepository) GetOrderByProductID(productID []string) (*[]model.TransactionLs, error) {
	var Order []model.TransactionLs

	rows, err := u.DB.Table("transactions_hydromart").Where("product_id IN ?", productID).Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order model.TransactionLs
		if err := u.DB.ScanRows(rows, &order); err != nil {
			return nil, err
		}
		Order = append(Order, order)
	}
	return &Order, nil
}

func (u *UserRepository) UpdateOrderStatus(order model.Transaction) error {
	result := u.DB.Table("transactions_hydromart").Where("transaction_id = ?", order.TransactionID).Updates(map[string]interface{}{
		"status": order.Status,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepository) GetOrderByID(orderID string) (model.Transaction, error) {
	var order model.Transaction
	result := u.DB.Table("transactions_hydromart").Where("transaction_id = ?", orderID).First(&order)
	if result.Error != nil {
		return model.Transaction{}, result.Error
	}
	return order, nil
}

func (u *UserRepository) GetUserByID(userID string) (*model.User, error) {
	var user model.User

	result := u.DB.Table("users_hydromart").Where("user_id", userID).First(&user)
	if result.Error != nil {
		return &model.User{}, result.Error
	}

	return &user, nil
}
