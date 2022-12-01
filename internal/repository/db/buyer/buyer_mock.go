package buyer

import (
	"errors"

	"github.com/hrz8/gokomodo-challenge/internal/model/entity"
	"github.com/stretchr/testify/mock"
)

type RepositoryBuyerMock struct {
	Mock mock.Mock
}

func (r *RepositoryBuyerMock) Create(data *entity.Buyer) (*entity.Buyer, error) {
	arguments := r.Mock.Called(data)
	if arguments.Get(0) == nil {
		return nil, errors.New("some error")
	}

	result := arguments.Get(0).(entity.Buyer)
	return &result, nil
}

func (r *RepositoryBuyerMock) FindById(id string) (*entity.Buyer, error) {
	arguments := r.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil, errors.New("some error")
	}

	result := arguments.Get(0).(entity.Buyer)
	return &result, nil
}

func (r *RepositoryBuyerMock) FindByEmail(email string) (*entity.Buyer, error) {
	arguments := r.Mock.Called(email)
	if arguments.Get(0) == nil {
		return nil, errors.New("some error")
	}

	result := arguments.Get(0).(entity.Buyer)
	return &result, nil
}
