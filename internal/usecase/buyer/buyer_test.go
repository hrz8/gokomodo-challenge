package buyer

import (
	"testing"

	"github.com/gofrs/uuid"
	"github.com/hrz8/gokomodo-challenge/internal/model/entity"
	"github.com/hrz8/gokomodo-challenge/internal/repository/db"
	"github.com/hrz8/gokomodo-challenge/internal/repository/db/buyer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	buyerRepositoryMock = &buyer.RepositoryBuyerMock{Mock: mock.Mock{}}
	dbRepositoryMock    = &db.DBRepositoryMock{Mock: mock.Mock{}}
	buyerUsecase        = NewUsecase(dbRepositoryMock)
)

func TestRegister(t *testing.T) {
	id, _ := uuid.NewV4()
	buyerMock := entity.Buyer{
		ID:               id,
		Email:            "some@email.com",
		Name:             "your name",
		Password:         "$2a$12$DKFBnEzTMuXah7.lMmS81u98iLjGcwGSvqTXBmXke3bT.ZbsY9.AK",
		RecipientAddress: "long address",
	}

	t.Run("Success login", func(t *testing.T) {
		buyerRepositoryMock.Mock.On("FindByEmail", "some@email.com").Return(buyerMock, nil)
		dbRepositoryMock.Mock.On("GetBuyerRepository").Return(buyerRepositoryMock)

		result, err := buyerUsecase.Login("some@email.com", "password")

		assert.Nil(t, err)
		assert.Equal(t, result.Email, buyerMock.Email)
	})

	t.Run("Wrong password login", func(t *testing.T) {
		buyerRepositoryMock.Mock.On("FindByEmail", "some@email.com").Return(nil, nil)
		dbRepositoryMock.Mock.On("GetBuyerRepository").Return(buyerRepositoryMock)

		result, err := buyerUsecase.Login("some@email.com", "wrongpassword")

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}
