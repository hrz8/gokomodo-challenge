package order

import (
	"github.com/hrz8/gokomodo-challenge/internal/model/entity"
	"gorm.io/gorm"
)

type (
	repository struct {
		Conn *gorm.DB
	}

	IRepositoryOrder interface {
		Create(data *entity.Order) (*entity.Order, error)
	}
)

func (r *repository) Create(data *entity.Order) (*entity.Order, error) {
	err := r.Conn.Debug().Create(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewRepository(conn *gorm.DB) IRepositoryOrder {
	return &repository{conn}
}
