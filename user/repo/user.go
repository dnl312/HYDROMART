package repo

import (
	"log"
	"user/model"

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

func (u *UserRepository) ValidateUser(userID string) (model.User, error) {
	var user model.User
	result := u.DB.Table("users_hydromart").Where("user_id = ? AND role= ? AND deposit>0", userID, "USER").First(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

func ( u *UserRepository) UpdateDeposit(order model.Transaction) error {
	result := u.DB.Table("users_hydromart").Where("user_id = ?", order.UserID).Update("deposit", gorm.Expr("deposit - ?", order.Amount))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepository) GetProductByID(productID string) (model.Product, error) {
	var product model.Product
	result := u.DB.Table("products_hydromart").Where("product_id = ?", productID).First(&product)
	if result.Error != nil {
		return model.Product{}, result.Error
	}
	return product, nil
}

func (u *UserRepository) CreateOrder(order model.Transaction) error {
	result := u.DB.Table("transactions_hydromart").Create(&order)
	if result != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepository) GetOrder(userID string, status string) (*[]model.TransactionLs, error) {
	var Order []model.TransactionLs
	if status != "" {
		log.Println("status not empty")
		rows, err := u.DB.Table("transactions_hydromart").Where("user_id = ? AND status = ?", userID, status).Rows()
	
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
	}else{
		rows, err := u.DB.Table("transactions_hydromart").Where("user_id = ?", userID).Rows()
	
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
	}

	return &Order, nil
}

func (u *UserRepository) GetOrderByID(orderID string) (model.Transaction, error) {
	var order model.Transaction
	result := u.DB.Table("transactions_hydromart").Where("transaction_id = ?", orderID).First(&order)
	if result.Error != nil {
		return model.Transaction{}, result.Error
	}
	return order, nil
}

func (u *UserRepository) UpdateOrder(order model.Transaction) error {
	result := u.DB.Table("transactions_hydromart").Where("transaction_id = ?", order.TransactionID).Updates(map[string]interface{}{
		"qty": order.Qty,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepository) DeleteOrder(orderID string) error {
	result := u.DB.Table("transactions_hydromart").Where("transaction_id = ? AND status = ? ", orderID, "ORDER CREATED").Delete(&model.Transaction{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
