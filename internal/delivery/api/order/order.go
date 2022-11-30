package order

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
	g.POST("", d.OrderProduct)
	g.PATCH("/:id/accept", d.Accept)
}

func (d *delivery) OrderProduct(ctx echo.Context) error {
	body := new([]dto.OrderProductRequest)

	token, err := auth.VerifyJWT(&ctx)
	if err != nil {
		return res.ErrorResponse(err).Send(ctx)
	}

	buyer, err := d.Usecase.Buyer.FindById(token.Sub)
	if err != nil {
		return res.ErrorBuilder(
			&res.ErrorConstant.Unauthorized,
			errors.New("you are not authorized"),
		).Send(ctx)
	}

	if buyer.ID.String() == "" || token.Aud != "buyer" {
		return res.ErrorBuilder(
			&res.ErrorConstant.Unauthorized,
			errors.New("you are not authorized"),
		).Send(ctx)
	}

	if err := ctx.Bind(body); err != nil {
		return res.ErrorBuilder(
			&res.ErrorConstant.InternalServerError,
			err,
		).Send(ctx)
	}

	data, err := d.Usecase.Order.OrderProduct(body, token.Sub)
	if err != nil {
		return res.ErrorResponse(err).Send(ctx)
	}

	return res.SuccessResponse(data).Send(ctx)
}

func (d *delivery) Accept(ctx echo.Context) error {
	id := ctx.Param("id")

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

	data, err := d.Usecase.Order.AcceptOrder(id)
	if err != nil {
		return res.ErrorResponse(err).Send(ctx)
	}

	return res.SuccessResponse(data).Send(ctx)
}

func NewDelivery(u *usecase.Usecase) *delivery {
	return &delivery{u}
}
