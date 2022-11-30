package seller

import (
	"fmt"

	"github.com/hrz8/gokomodo-challenge/internal/model/entity"
	"gorm.io/gorm"
)

type (
	repository struct {
		Conn *gorm.DB
	}

	IRepositorySeller interface {
		FindByEmail(email string) (*entity.Seller, error)
	}
)

func (r *repository) FindByEmail(email string) (*entity.Seller, error) {
	fmt.Println("this is repository", r)
	result := &entity.Seller{}
	return result, nil
}

func NewRepository(conn *gorm.DB) IRepositorySeller {
	return &repository{
		Conn: conn,
	}
}
