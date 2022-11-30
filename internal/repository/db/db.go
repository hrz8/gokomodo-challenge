package db

import (
	"github.com/hrz8/gokomodo-challenge/internal/repository/db/buyer"
	"github.com/hrz8/gokomodo-challenge/internal/repository/db/order"
	"github.com/hrz8/gokomodo-challenge/internal/repository/db/product"
	"github.com/hrz8/gokomodo-challenge/internal/repository/db/seller"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB

	Buyer   buyer.IRepositoryBuyer
	Seller  seller.IRepositorySeller
	Product product.IRepositoryProduct
	Order   order.IRepositoryOrder
}

func NewRepository(conn *gorm.DB) *Repository {
	return &Repository{
		DB: conn,

		Buyer:   buyer.NewRepository(conn),
		Seller:  seller.NewRepository(conn),
		Product: product.NewRepository(conn),
		Order:   order.NewRepository(conn),
	}
}
