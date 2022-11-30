package product

import (
	"github.com/hrz8/gokomodo-challenge/internal/model/entity"
	"github.com/hrz8/gokomodo-challenge/pkg/util/pagination"
	"gorm.io/gorm"
)

type (
	repository struct {
		Conn *gorm.DB
	}

	IRepositoryProduct interface {
		Create(data *entity.Product) (*entity.Product, error)
		List(page uint16, limit uint16) (*[]entity.Product, error)
	}
)

func (r *repository) Create(data *entity.Product) (*entity.Product, error) {
	if err := r.Conn.Debug().Create(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *repository) List(page uint16, limit uint16) (*[]entity.Product, error) {
	result := []entity.Product{}

	if err := r.Conn.Debug().
		Limit(int(limit)).
		Offset(pagination.GetOffset(int(page), int(limit))).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func NewRepository(conn *gorm.DB) IRepositoryProduct {
	return &repository{conn}
}
