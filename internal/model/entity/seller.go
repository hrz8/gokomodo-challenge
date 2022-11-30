package entity

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Seller struct {
	gorm.Model
	ID            uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	Email         string    `gorm:"column:email;index:idx_seller_email;unique;not null" json:"email" validate:"required"`
	Name          string    `gorm:"column:name;not null" json:"name" validate:"required"`
	Password      string    `gorm:"column:password;not null" json:"-" validate:"required"`
	PickupAddress string    `gorm:"column:pickup_address" json:"recipient_address" validate:"required"`
	// timestamp
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}
