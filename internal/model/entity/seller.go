package entity

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Seller struct {
	ID               uuid.UUID      `gorm:"column:id;primaryKey" json:"id"`
	Email            string         `gorm:"column:email;index:idx_email;unique;not null" json:"email" validate:"required"`
	Name             string         `gorm:"column:name;not null" json:"name" validate:"required"`
	Password         string         `gorm:"column:password;not null" json:"-" validate:"required"`
	RecipientAddress string         `gorm:"column:recipient_address" json:"recipient_address" validate:"required"`
	CreatedAt        time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt        time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}