package dto

import (
	"github.com/gofrs/uuid"
	"github.com/hrz8/gokomodo-challenge/internal/model/entity"
)

type (
	OrderProductRequest struct {
		ProductID string `json:"product_id" validate:"required"`
		Quantity  uint16 `json:"quantity" validate:"required"`
	}

	OrderedProduct struct {
		ProductID uuid.UUID `json:"product_id"`
		Quantity  uint16    `json:"quantity"`
		Price     uint16    `json:"price"`
	}

	OrderProductResponse struct {
		OrderID    string           `json:"order_id"`
		Products   []OrderedProduct `json:"products"`
		TotalPrice uint16           `json:"total_price"`
	}

	OrderItemResponse struct {
		ProductID   uuid.UUID `json:"product_id"`
		ProductName string    `json:"product_name"`
		Price       uint16    `json:"price"`
		Quantity    uint16    `json:"quantity"`
	}

	OrderDetailResponse struct {
		ID            uuid.UUID           `json:"id"`
		Status        entity.OrderStatus  `json:"status"`
		TotalPrice    uint16              `json:"total_price"`
		BuyerName     string              `json:"buyer_name"`
		BuyerAddress  string              `json:"buyer_address"`
		SellerName    string              `json:"seller_name"`
		SellerAddress string              `json:"seller_address"`
		Items         []OrderItemResponse `json:"items"`
	}
)
