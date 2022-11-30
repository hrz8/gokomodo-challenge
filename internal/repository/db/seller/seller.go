package seller

import (
	"github.com/hrz8/gokomodo-challenge/internal/model/entity"
	"gorm.io/gorm"
)

type (
	repository struct {
		Conn *gorm.DB
	}

	IRepositorySeller interface {
		Create(data *entity.Seller) (*entity.Seller, error)
		FindById(id string) (*entity.Seller, error)
		FindByEmail(email string) (*entity.Seller, error)
	}
)

func (r *repository) Create(data *entity.Seller) (*entity.Seller, error) {
	if err := r.Conn.Debug().Create(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *repository) FindById(id string) (*entity.Seller, error) {
	result := new(entity.Seller)
	if err := r.Conn.Debug().Where("`id` = ?", id).First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) FindByEmail(email string) (*entity.Seller, error) {
	result := new(entity.Seller)
	if err := r.Conn.Debug().
		Where("`email` = ?", email).
		First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func NewRepository(conn *gorm.DB) IRepositorySeller {
	return &repository{conn}
}
