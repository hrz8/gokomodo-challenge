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
)
