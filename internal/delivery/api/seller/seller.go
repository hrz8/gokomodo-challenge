package seller

import (
	"github.com/hrz8/gokomodo-challenge/internal/model/dto"
	"github.com/hrz8/gokomodo-challenge/internal/usecase"
	res "github.com/hrz8/gokomodo-challenge/pkg/util/response"
	"github.com/labstack/echo/v4"
)

type delivery struct {
	Usecase *usecase.Usecase
}

func (d *delivery) Route(g *echo.Group) {
	g.POST("/register", d.Register)
	g.POST("/login", d.Login)
}

func (d *delivery) Register(ctx echo.Context) error {
	body := new(dto.RegisterRequest)

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

	data, err := d.Usecase.SellerRepository.Register(body)
	if err != nil {
		return res.ErrorResponse(err).Send(ctx)
	}

	return res.SuccessResponse(dto.RegisterResponse{
		ID:      data.ID,
		Name:    data.Name,
		Email:   data.Email,
		Address: data.PickupAddress,
	}).Send(ctx)
}

func (d *delivery) Login(ctx echo.Context) error {
	body := new(dto.LoginRequest)

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

	result, err := d.Usecase.SellerRepository.Login(body.Email, body.Password)
	if err != nil {
		return res.ErrorResponse(err).Send(ctx)
	}

	token, err := d.Usecase.SellerRepository.GenerateToken(result.ID.String())
	if err != nil {
		return res.ErrorResponse(err).Send(ctx)
	}

	return res.SuccessResponse(dto.LoginResponse{
		Email: result.Email,
		Token: token,
	}).Send(ctx)
}

func NewDelivery(u *usecase.Usecase) *delivery {
	return &delivery{u}
}
