package http

import (
	"fmt"
	"net/http"

	"github.com/hrz8/gokomodo-challenge/internal/delivery/api"
	"github.com/hrz8/gokomodo-challenge/internal/repository/db"
	"github.com/hrz8/gokomodo-challenge/internal/usecase"
	"github.com/hrz8/gokomodo-challenge/pkg/util/validator"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type (
	driver struct {
		e *echo.Echo
	}

	IDriverHttp interface {
		Start(conn *gorm.DB) error
	}
)

func (d *driver) Start(conn *gorm.DB) error {
	d.e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome!")
	})

	// build
	repository := db.NewRepository(conn)
	usecase := usecase.NewUsecase(repository)

	// delivery
	api.NewDelivery(repository, usecase).Start(d.e)

	return d.e.Start(fmt.Sprintf(":%d", 3000))
}

func NewDriver() IDriverHttp {
	e := echo.New()

	// middleware
	e.Validator = &validator.CustomValidator{
		Validator: validator.NewValidator(),
	}

	return &driver{e}
}
