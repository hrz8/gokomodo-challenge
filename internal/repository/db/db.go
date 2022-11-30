package db

import (
	"github.com/hrz8/gokomodo-challenge/internal/repository/db/buyer"
	"github.com/hrz8/gokomodo-challenge/internal/repository/db/seller"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB

	BuyerRepository  buyer.IRepositoryBuyer
	SellerRepository seller.IRepositorySeller
}

func NewRepository(conn *gorm.DB) *Repository {
	return &Repository{
		DB: conn,

		BuyerRepository:  buyer.NewRepository(conn),
		SellerRepository: seller.NewRepository(conn),
	}
}
