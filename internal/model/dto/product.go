package dto

import "github.com/gofrs/uuid"

type (
	AddProductRequest struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
		Price       uint16 `json:"price" validate:"required"`
	}

	AddProductResponse struct {
		ID          uuid.UUID `json:"id"`
		SellerID    uuid.UUID `json:"seller_id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Price       uint16    `json:"price"`
	}

	ProductDetailResponse struct {
		ID          uuid.UUID `json:"id"`
		SellerID    uuid.UUID `json:"seller_id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Price       uint16    `json:"price"`
	}
)
