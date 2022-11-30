package db

import (
	"github.com/hrz8/gokomodo-challenge/internal/repository/db/buyer"
	"github.com/hrz8/gokomodo-challenge/internal/repository/db/seller"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB

	Buyer  buyer.IRepositoryBuyer
	Seller seller.IRepositorySeller
}

func NewRepository(conn *gorm.DB) *Repository {
	return &Repository{
		DB: conn,

		Buyer:  buyer.NewRepository(conn),
		Seller: seller.NewRepository(conn),
	}
}
