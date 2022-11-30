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
		Accept(id string) (entity.OrderStatus, error)
	}
)

func (r *repository) Create(data *entity.Order) (*entity.Order, error) {
	err := r.Conn.Debug().Create(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *repository) FindById(id string) (*entity.Order, error) {
	result := new(entity.Order)
	if err := r.Conn.Debug().Where("`id` = ?", id).First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) Accept(id string) (entity.OrderStatus, error) {
	_, err := r.FindById(id)
	if err != nil {
		return "", err
	}

	if err := r.Conn.Model(&entity.Order{}).Where("`id` = ?", id).Update("status", entity.ACCEPTED).Error; err != nil {
		return "", err
	}

	return entity.ACCEPTED, nil
}

func NewRepository(conn *gorm.DB) IRepositoryOrder {
	return &repository{conn}
}
