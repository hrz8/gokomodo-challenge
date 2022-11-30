package order

import (
	"github.com/gofrs/uuid"
	"github.com/hrz8/gokomodo-challenge/internal/model/dto"
	"github.com/hrz8/gokomodo-challenge/internal/model/entity"
	"github.com/hrz8/gokomodo-challenge/internal/repository/db"
	res "github.com/hrz8/gokomodo-challenge/pkg/util/response"
	"github.com/pkg/errors"
)

type (
	usecase struct {
		Repository *db.Repository
	}

	IUsecaseOrder interface {
		OrderProduct(body *[]dto.OrderProductRequest, sub string) (*dto.OrderProductResponse, error)
		AcceptOrder(id string) (map[string]string, error)
	}
)

func (u *usecase) OrderProduct(body *[]dto.OrderProductRequest, sub string) (*dto.OrderProductResponse, error) {
	id, _ := uuid.NewV4()
	sellerIds := []string{}
	products := []dto.OrderedProduct{}

	buyer, err := u.Repository.Buyer.FindById(sub)
	if err != nil {
		return nil, res.ErrorBuilder(
			&res.ErrorConstant.InternalServerError,
			err,
		)
	}

	data := &entity.Order{
		ID:            id,
		TotalPrice:    0,
		Item:          []entity.Item{},
		Status:        entity.PENDING,
		BuyerID:       uuid.FromStringOrNil(sub),
		SourceAddress: buyer.RecipientAddress,
	}

	for i, item := range *body {
		if item.ProductID == "" || item.Quantity == 0 {
			return nil, res.ErrorBuilder(
				&res.ErrorConstant.BadRequest,
				errors.New(""),
				"missing product id",
			)
		}
		product, err := u.Repository.Product.FindById(item.ProductID)
		if err != nil {
			return nil, res.ErrorBuilder(
				&res.ErrorConstant.NotFound,
				err,
				"cannot found product id",
			)
		}

		sellerId := product.SellerID.String()

		if i == 0 {
			seller, err := u.Repository.Seller.FindById(sellerId)
			if err != nil {
				return nil, res.ErrorBuilder(
					&res.ErrorConstant.InternalServerError,
					err,
				)
			}

			data.SellerID = seller.ID
			data.DestinationAddress = seller.PickupAddress

			sellerIds = append(sellerIds, sellerId)
		}

		if i != 0 && sellerIds[0] != sellerId {
			return nil, res.ErrorBuilder(
				&res.ErrorConstant.BadRequest,
				errors.New(""),
				"cannot create order from multiple seller",
			)
		}

		data.TotalPrice += product.Price * item.Quantity
		data.Item = append(data.Item, entity.Item{
			ProductID:   product.ID,
			ProductName: product.ProductName,
			Price:       product.Price,
			Quantity:    item.Quantity,
		})

		products = append(products, dto.OrderedProduct{
			ProductID: product.ID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		})
	}

	result, err := u.Repository.Order.Create(data)
	if err != nil {
		return nil, res.ErrorBuilder(
			&res.ErrorConstant.InternalServerError,
			err,
		)
	}

	return &dto.OrderProductResponse{
		OrderID:    result.ID.String(),
		Products:   products,
		TotalPrice: result.TotalPrice,
	}, nil
}

func (u *usecase) AcceptOrder(id string) (map[string]string, error) {
	result := map[string]string{}
	status, err := u.Repository.Order.Accept(id)
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

	result["status"] = string(status)

	return result, nil
}

func NewUsecase(r *db.Repository) IUsecaseOrder {
	return &usecase{r}
}
