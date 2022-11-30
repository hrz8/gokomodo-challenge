package product

import (
	"errors"

	"github.com/hrz8/gokomodo-challenge/internal/model/dto"
	"github.com/hrz8/gokomodo-challenge/internal/usecase"
	"github.com/hrz8/gokomodo-challenge/pkg/util/auth"
	res "github.com/hrz8/gokomodo-challenge/pkg/util/response"
	"github.com/labstack/echo/v4"
)

type delivery struct {
	Usecase *usecase.Usecase
}

func (d *delivery) Route(g *echo.Group) {
	g.POST("", d.AddNewProduct)
}

func (d *delivery) AddNewProduct(ctx echo.Context) error {
	token, err := auth.VerifyJWT(&ctx)
	if err != nil {
		return res.ErrorResponse(err).Send(ctx)
	}

	seller, err := d.Usecase.Seller.FindById(token.Sub)
	if err != nil {
		return res.ErrorBuilder(
			&res.ErrorConstant.Unauthorized,
			errors.New("you are not authorized"),
		).Send(ctx)
	}

	if seller.ID.String() == "" || token.Aud != "seller" {
		return res.ErrorBuilder(
			&res.ErrorConstant.Unauthorized,
			errors.New("you are not authorized"),
		).Send(ctx)
	}

	// main
	body := new(dto.AddProductRequest)

	if err := ctx.Bind(body); err != nil {
		return res.ErrorBuilder(
			&res.ErrorConstant.InternalServerError,
			err,
		).Send(ctx)
	}

	if err := ctx.Validate(body); err != nil {
		return res.ErrorBuilder(
			&res.ErrorConstant.BadRequest,
			err,
		).Send(ctx)
	}

	data, err := d.Usecase.Product.AddProduct(body, seller)
	if err != nil {
		return res.ErrorResponse(err).Send(ctx)
	}

	return res.SuccessResponse(data).Send(ctx)
}

func NewDelivery(u *usecase.Usecase) *delivery {
	return &delivery{u}
}
