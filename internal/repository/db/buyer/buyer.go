package buyer

import (
	"github.com/hrz8/gokomodo-challenge/internal/model/entity"
	"gorm.io/gorm"
)

type (
	repository struct {
		Conn *gorm.DB
	}

	IRepositoryBuyer interface {
		Create(data *entity.Buyer) (*entity.Buyer, error)
		FindById(string string) (*entity.Buyer, error)
		FindByEmail(email string) (*entity.Buyer, error)
	}
)

func (r *repository) Create(data *entity.Buyer) (*entity.Buyer, error) {
	err := r.Conn.Debug().Create(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *repository) FindById(id string) (*entity.Buyer, error) {
	result := new(entity.Buyer)
	err := r.Conn.Debug().Where("`id` = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) FindByEmail(email string) (*entity.Buyer, error) {
	result := new(entity.Buyer)
	err := r.Conn.Debug().Where("`email` = ?", email).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewRepository(conn *gorm.DB) IRepositoryBuyer {
	return &repository{conn}
}
