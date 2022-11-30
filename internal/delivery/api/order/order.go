package order

import (
	"errors"

	"github.com/hrz8/gokomodo-challenge/internal/model/dto"
	"github.com/hrz8/gokomodo-challenge/internal/usecase"
	"github.com/hrz8/gokomodo-challenge/pkg/util/auth"
	"github.com/hrz8/gokomodo-challenge/pkg/util/parser"
	res "github.com/hrz8/gokomodo-challenge/pkg/util/response"
	"github.com/labstack/echo/v4"
)

type delivery struct {
	Usecase *usecase.Usecase
}

func (d *delivery) Route(g *echo.Group) {
	g.GET("", d.ListOrders)
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

func (d *delivery) ListOrders(ctx echo.Context) error {
	token, err := auth.VerifyJWT(&ctx)
	if err != nil {
		return res.ErrorResponse(err).Send(ctx)
	}

	isBuyer := false

	buyer, _ := d.Usecase.Buyer.FindById(token.Sub)
	seller, _ := d.Usecase.Seller.FindById(token.Sub)

	if buyer != nil && seller == nil {
		isBuyer = true
	}

	queryParams := ctx.QueryParams()
	qpPage, pageExist := queryParams["page"]
	qpLimit, limitExists := queryParams["limit"]

	if !pageExist || !limitExists {
		qpPage = []string{"1"}
		qpLimit = []string{"10"}
	}

	page := uint16(parser.ParseStringToInt(qpPage[0]))
	limit := uint16(parser.ParseStringToInt(qpLimit[0]))

	data, err := d.Usecase.Order.ListOrders(page, limit, isBuyer, token.Sub)
	if err != nil {
		return res.ErrorResponse(err).Send(ctx)
	}

	return res.SuccessResponse(data).Send(ctx)
}

func NewDelivery(u *usecase.Usecase) *delivery {
	return &delivery{u}
}
