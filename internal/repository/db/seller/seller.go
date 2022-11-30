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
		FindByEmail(email string) (*entity.Seller, error)
	}
)

func (r *repository) Create(data *entity.Seller) (*entity.Seller, error) {
	err := r.Conn.Create(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *repository) FindByEmail(email string) (*entity.Seller, error) {
	result := new(entity.Seller)
	err := r.Conn.Debug().Where("`email` = ?", email).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewRepository(conn *gorm.DB) IRepositorySeller {
	return &repository{conn}
}
