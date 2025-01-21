package repo

import (
	"errors"
	"merchant/model"

	"github.com/stretchr/testify/mock"
)

// define mock user repository
type MerchantRepositoryMock struct {
	Mock mock.Mock
}

func (m *MerchantRepositoryMock) GetAllProduct(id string) (*[]model.Product, error) {
	res := m.Mock.Called(id)

	if res.Get(0) == nil {
		return nil, errors.New("user id not found")
	}

	product := res.Get(0).(model.Product)
	productList := []model.Product{
		product,
	}
	return &productList, nil
}

// AddUser inserts user to database mock
func (m *MerchantRepositoryMock) AddProduct(userPtr *model.Product) error {
	res := m.Mock.Called(userPtr)

	if res.Get(0) == nil {
		return errors.New("adding product failed")
	}

	return nil
}

// UpdateProduct updates user in database mock
func (m *MerchantRepositoryMock) UpdateProduct(userPtr *model.Product) error {
	res := m.Mock.Called(userPtr)

	if res.Get(0) == nil {
		return errors.New("updating product failed")
	}

	return nil
}

// UpdateProduct updates user in database mock
func (m *MerchantRepositoryMock) DeleteProduct(productID, merchantID string) error {
	res := m.Mock.Called(productID, merchantID)

	if res.Get(0) == nil {
		return errors.New("delete product failed")
	}

	return nil
}
