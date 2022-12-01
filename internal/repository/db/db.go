package db

import (
	"github.com/hrz8/gokomodo-challenge/internal/repository/db/buyer"
	"github.com/hrz8/gokomodo-challenge/internal/repository/db/order"
	"github.com/hrz8/gokomodo-challenge/internal/repository/db/product"
	"github.com/hrz8/gokomodo-challenge/internal/repository/db/seller"
	"gorm.io/gorm"
)

type (
	Repository struct {
		DB *gorm.DB

		Buyer   buyer.IRepositoryBuyer
		Seller  seller.IRepositorySeller
		Product product.IRepositoryProduct
		Order   order.IRepositoryOrder
	}

	IDBRepository interface {
		GetConn() *gorm.DB
		GetBuyerRepository() buyer.IRepositoryBuyer
		GetSellerRepository() seller.IRepositorySeller
		GetProductRepository() product.IRepositoryProduct
		GetOrderRepository() order.IRepositoryOrder
	}
)

func (r *Repository) GetConn() *gorm.DB {
	return r.DB
}

func (r *Repository) GetBuyerRepository() buyer.IRepositoryBuyer {
	return r.Buyer
}

func (r *Repository) GetSellerRepository() seller.IRepositorySeller {
	return r.Seller
}

func (r *Repository) GetProductRepository() product.IRepositoryProduct {
	return r.Product
}

func (r *Repository) GetOrderRepository() order.IRepositoryOrder {
	return r.Order
}

func NewRepository(conn *gorm.DB) IDBRepository {
	return &Repository{
		DB: conn,

		Buyer:   buyer.NewRepository(conn),
		Seller:  seller.NewRepository(conn),
		Product: product.NewRepository(conn),
		Order:   order.NewRepository(conn),
	}
}
