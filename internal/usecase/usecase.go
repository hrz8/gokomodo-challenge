package usecase

import (
	"github.com/hrz8/gokomodo-challenge/internal/repository/db"
	"github.com/hrz8/gokomodo-challenge/internal/usecase/buyer"
	"github.com/hrz8/gokomodo-challenge/internal/usecase/seller"
)

type Usecase struct {
	DBRepository *db.Repository

	BuyerUsecase     buyer.IUsecaseBuyer
	SellerRepository seller.IUsecaseSeller
}

func NewUsecase(dbRepository *db.Repository) *Usecase {
	return &Usecase{
		DBRepository: dbRepository,

		BuyerUsecase:     buyer.NewUsecase(dbRepository),
		SellerRepository: seller.NewUsecase(dbRepository),
	}
}
