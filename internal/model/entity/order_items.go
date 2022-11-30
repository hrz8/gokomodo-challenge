package entity

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type (
	OrderItem struct {
		gorm.Model
		ID       uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
		Price    uint16    `gorm:"column:price;not null" json:"price" validate:"required"`
		Quantity uint16    `gorm:"column:quantity;not null" json:"quantity" validate:"required"`
		// has one - required
		OrderID   uuid.UUID `gorm:"size:40" json:"-"`
		Order     Order     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"order,omitempty"`
		ProductID uuid.UUID `gorm:"size:40" json:"-"`
		Product   Product   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product,omitempty"`
		// timestamp
		CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`
		UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`
		DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	}

	Item struct {
		gorm.Model
		ProductID   uuid.UUID `json:"product_id"`
		ProductName string    `json:"product_name"`
		Price       uint16    `json:"price"`
		Quantity    uint16    `json:"quantity"`
	}
)
