package api

import (
	"github.com/hrz8/gokomodo-challenge/internal/delivery/api/buyer"
	"github.com/hrz8/gokomodo-challenge/internal/delivery/api/order"
	"github.com/hrz8/gokomodo-challenge/internal/delivery/api/product"
	"github.com/hrz8/gokomodo-challenge/internal/delivery/api/seller"
	"github.com/hrz8/gokomodo-challenge/internal/repository/db"
	"github.com/hrz8/gokomodo-challenge/internal/usecase"
	"github.com/labstack/echo/v4"
)

type (
	delivery struct {
		Repository db.IDBRepository
		Usecase    *usecase.Usecase
	}

	IDeliveryApi interface {
		Start(e *echo.Echo)
	}
)

func (d *delivery) Start(e *echo.Echo) {
	buyer.NewDelivery(d.Usecase).Route(e.Group("buyer"))
	seller.NewDelivery(d.Usecase).Route(e.Group("seller"))
	product.NewDelivery(d.Usecase).Route(e.Group("products"))
	order.NewDelivery(d.Usecase).Route(e.Group("orders"))
}

func NewDelivery(
	r db.IDBRepository,
	u *usecase.Usecase,
) IDeliveryApi {
	return &delivery{r, u}
}
