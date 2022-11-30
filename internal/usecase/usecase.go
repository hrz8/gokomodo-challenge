package usecase

import (
	"github.com/hrz8/gokomodo-challenge/internal/repository/db"
	"github.com/hrz8/gokomodo-challenge/internal/usecase/buyer"
	"github.com/hrz8/gokomodo-challenge/internal/usecase/product"
	"github.com/hrz8/gokomodo-challenge/internal/usecase/seller"
)

type Usecase struct {
	DBRepository *db.Repository

	Buyer   buyer.IUsecaseBuyer
	Seller  seller.IUsecaseSeller
	Product product.IUsecaseSeller
}

func NewUsecase(dbRepository *db.Repository) *Usecase {
	return &Usecase{
		DBRepository: dbRepository,

		Buyer:   buyer.NewUsecase(dbRepository),
		Seller:  seller.NewUsecase(dbRepository),
		Product: product.NewUsecase(dbRepository),
	}
}
