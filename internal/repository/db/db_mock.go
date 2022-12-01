package db

import (
	"github.com/hrz8/gokomodo-challenge/internal/repository/db/buyer"
	"github.com/hrz8/gokomodo-challenge/internal/repository/db/order"
	"github.com/hrz8/gokomodo-challenge/internal/repository/db/product"
	"github.com/hrz8/gokomodo-challenge/internal/repository/db/seller"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type DBRepositoryMock struct {
	Mock mock.Mock
}

func (r *DBRepositoryMock) GetConn() *gorm.DB {
	arguments := r.Mock.Called()
	if arguments.Get(0) == nil {
		return nil
	}

	result := arguments.Get(0).(gorm.DB)
	return &result
}

func (r *DBRepositoryMock) GetBuyerRepository() buyer.IRepositoryBuyer {
	arguments := r.Mock.Called()
	if arguments.Get(0) == nil {
		return nil
	}

	result := arguments.Get(0).(buyer.IRepositoryBuyer)
	return result
}

func (r *DBRepositoryMock) GetSellerRepository() seller.IRepositorySeller {
	arguments := r.Mock.Called()
	if arguments.Get(0) == nil {
		return nil
	}

	result := arguments.Get(0).(seller.IRepositorySeller)
	return result
}

func (r *DBRepositoryMock) GetProductRepository() product.IRepositoryProduct {
	arguments := r.Mock.Called()
	if arguments.Get(0) == nil {
		return nil
	}

	result := arguments.Get(0).(product.IRepositoryProduct)
	return result
}

func (r *DBRepositoryMock) GetOrderRepository() order.IRepositoryOrder {
	arguments := r.Mock.Called()
	if arguments.Get(0) == nil {
		return nil
	}

	result := arguments.Get(0).(order.IRepositoryOrder)
	return result
}
