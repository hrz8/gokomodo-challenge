package entity

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	ProductName string    `gorm:"column:product_name;not null" json:"product_name" validate:"required"`
	Description string    `gorm:"column:description" json:"description" validate:"required"`
	Price       uint16    `gorm:"column:price" json:"price" validate:"required"`
	// has one - required
	SellerID uuid.UUID `gorm:"size:40" json:"-"`
	Seller   Seller    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"seller,omitempty"`
	// timestamp
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}
