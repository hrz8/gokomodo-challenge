package buyer

import (
	"github.com/hrz8/gokomodo-challenge/internal/model/dto"
	"github.com/hrz8/gokomodo-challenge/internal/repository/db"
)

type (
	usecase struct {
		Repository *db.Repository
	}
	IUsecaseBuyer interface {
		Login(body dto.LoginRequest) (*dto.LoginResponse, error)
	}
)

func (u *usecase) Login(body dto.LoginRequest) (*dto.LoginResponse, error) {
	data, err := u.Repository.BuyerRepository.FindByEmail(body.Email)
	if err != nil {
		return nil, nil
	}

	result := &dto.LoginResponse{
		Email: data.Email,
		Token: "token",
	}

	return result, nil
}

func NewUsecase(r *db.Repository) IUsecaseBuyer {
	return &usecase{r}
}
