package seller

import (
	"github.com/hrz8/gokomodo-challenge/internal/model/dto"
	"github.com/hrz8/gokomodo-challenge/internal/repository/db"
)

type (
	usecase struct {
		Repository *db.Repository
	}

	IUsecaseSeller interface {
		Login(body dto.LoginRequest) (*dto.LoginResponse, error)
	}
)

func (u *usecase) Login(body dto.LoginRequest) (*dto.LoginResponse, error) {
	data, err := u.Repository.SellerRepository.FindByEmail(body.Email)
	if err != nil {
		return nil, nil
	}

	result := &dto.LoginResponse{
		Email: data.Email,
		Token: "token",
	}

	return result, nil
}

func NewUsecase(r *db.Repository) IUsecaseSeller {
	return &usecase{
		Repository: r,
	}
}
