package dto

import "github.com/gofrs/uuid"

type (
	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	LoginResponse struct {
		Email string `json:"email"`
		Token string `json:"token"`
	}

	RegisterRequest struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
		Address  string `json:"address" validate:"required"`
	}

	RegisterResponse struct {
		ID      uuid.UUID `json:"id"`
		Name    string    `json:"name"`
		Email   string    `json:"email"`
		Address string    `json:"address"`
	}

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
)
