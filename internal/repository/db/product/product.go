package product

import (
	"github.com/hrz8/gokomodo-challenge/internal/model/entity"
	"gorm.io/gorm"
)

type (
	repository struct {
		Conn *gorm.DB
	}

	IRepositoryProduct interface {
		Create(data *entity.Product) (*entity.Product, error)
	}
)

func (r *repository) Create(data *entity.Product) (*entity.Product, error) {
	err := r.Conn.Debug().Create(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewRepository(conn *gorm.DB) IRepositoryProduct {
	return &repository{conn}
}
