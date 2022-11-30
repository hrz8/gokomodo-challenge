package product

import (
	"github.com/gofrs/uuid"
	"github.com/hrz8/gokomodo-challenge/internal/model/dto"
	"github.com/hrz8/gokomodo-challenge/internal/model/entity"
	"github.com/hrz8/gokomodo-challenge/internal/repository/db"
	res "github.com/hrz8/gokomodo-challenge/pkg/util/response"
)

type (
	usecase struct {
		Repository *db.Repository
	}

	IUsecaseSeller interface {
		AddProduct(body *dto.AddProductRequest, seller *entity.Seller) (*dto.AddProductResponse, error)
	}
)

func (u *usecase) AddProduct(body *dto.AddProductRequest, seller *entity.Seller) (*dto.AddProductResponse, error) {
	id, _ := uuid.NewV4()

	data := &entity.Product{
		ID:          id,
		ProductName: body.Name,
		Description: body.Description,
		Price:       body.Price,
		SellerID:    seller.ID,
	}

	result, err := u.Repository.Product.Create(data)
	if err != nil {
		return nil, res.ErrorBuilder(
			&res.ErrorConstant.InternalServerError,
			err,
		)
	}

	return &dto.AddProductResponse{
		ID:          result.ID,
		SellerID:    seller.ID,
		Name:        result.ProductName,
		Description: result.Description,
		Price:       result.Price,
	}, nil
}

func NewUsecase(r *db.Repository) IUsecaseSeller {
	return &usecase{r}
}
