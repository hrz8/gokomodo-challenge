package entity

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type OrderStatus string

const (
	PENDING  OrderStatus = "PENDING"
	ACCEPTED OrderStatus = "ACCEPTED"
)

type Order struct {
	gorm.Model
	ID                 uuid.UUID   `gorm:"column:id;primaryKey" json:"id"`
	SourceAddress      string      `gorm:"column:source_address;not null" json:"source_address" validate:"required"`
	DestinationAddress string      `gorm:"column:destination_address;not null" json:"destination_address" validate:"required"`
	TotalPrice         uint16      `gorm:"column:total_price;not null" json:"total_price" validate:"required"`
	Status             OrderStatus `gorm:"column:status;not null" json:"status" validate:"required"`
	// has one - required
	BuyerID  uuid.UUID `gorm:"size:40" json:"-"`
	Buyer    Buyer     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"buyer,omitempty"`
	SellerID uuid.UUID `gorm:"size:40" json:"-"`
	Seller   Seller    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"seller,omitempty"`
	// timestamp
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	// etc
	Item []Item `gorm:"many2many:order_items;foreignKey:ID;joinForeignKey:OrderID;References:ID;JoinReferences:ProductID" json:"items,omitempty"`
}
