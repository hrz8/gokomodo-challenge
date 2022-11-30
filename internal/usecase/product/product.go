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
		FindById(id string) (*entity.Product, error)
		ListProducts(page uint16, limit uint16) (*[]dto.ProductDetailResponse, error)
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

func (u *usecase) FindById(id string) (*entity.Product, error) {
	result, err := u.Repository.Product.FindById(id)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, res.ErrorBuilder(
				&res.ErrorConstant.NotFound,
				err,
			)
		}
		return nil, res.ErrorBuilder(
			&res.ErrorConstant.InternalServerError,
			err,
		)
	}

	return result, nil
}

func (u *usecase) ListProducts(page uint16, limit uint16) (*[]dto.ProductDetailResponse, error) {
	result := []dto.ProductDetailResponse{}

	data, err := u.Repository.Product.List(page, limit)
	if err != nil {
		return nil, res.ErrorBuilder(
			&res.ErrorConstant.InternalServerError,
			err,
		)
	}

	for _, d := range *data {
		result = append(result, dto.ProductDetailResponse{
			ID:          d.ID,
			SellerID:    d.SellerID,
			Name:        d.ProductName,
			Description: d.Description,
			Price:       d.Price,
		})
	}

	return &result, nil
}

func NewUsecase(r *db.Repository) IUsecaseSeller {
	return &usecase{r}
}
