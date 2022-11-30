package buyer

import (
	"fmt"

	"github.com/hrz8/gokomodo-challenge/internal/model/entity"
	"gorm.io/gorm"
)

type (
	repository struct {
		Conn *gorm.DB
	}

	IRepositoryBuyer interface {
		FindByEmail(email string) (*entity.Buyer, error)
	}
)

func (r *repository) FindByEmail(email string) (*entity.Buyer, error) {
	fmt.Println("this is repository", r)
	result := &entity.Buyer{}
	return result, nil
}

func NewRepository(conn *gorm.DB) IRepositoryBuyer {
	return &repository{
		Conn: conn,
	}
}
