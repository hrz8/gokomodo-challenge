package usecase

import (
	"github.com/hrz8/gokomodo-challenge/internal/repository/db"
	"github.com/hrz8/gokomodo-challenge/internal/usecase/buyer"
	"github.com/hrz8/gokomodo-challenge/internal/usecase/order"
	"github.com/hrz8/gokomodo-challenge/internal/usecase/product"
	"github.com/hrz8/gokomodo-challenge/internal/usecase/seller"
)

type Usecase struct {
	DBRepository db.IDBRepository

	Buyer   buyer.IUsecaseBuyer
	Seller  seller.IUsecaseSeller
	Product product.IUsecaseSeller
	Order   order.IUsecaseOrder
}

func NewUsecase(dbRepository db.IDBRepository) *Usecase {
	return &Usecase{
		DBRepository: dbRepository,

		Buyer:   buyer.NewUsecase(dbRepository),
		Seller:  seller.NewUsecase(dbRepository),
		Product: product.NewUsecase(dbRepository),
		Order:   order.NewUsecase(dbRepository),
	}
}
