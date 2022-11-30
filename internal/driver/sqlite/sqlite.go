package sqlite

import (
	"log"

	"github.com/hrz8/gokomodo-challenge/internal/model/entity"
	sqliteDriver "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type (
	driver struct {
		File string
	}

	IDriverSqlite interface {
		Start() *gorm.DB
	}
)

func (d *driver) Start() *gorm.DB {
	db, err := gorm.Open(sqliteDriver.Open(d.File), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open connection to database!")
	}

	db.AutoMigrate(&entity.Buyer{})
	db.AutoMigrate(&entity.Seller{})
	db.AutoMigrate(&entity.Product{})
	db.AutoMigrate(&entity.Order{})
	db.AutoMigrate(&entity.OrderItem{})

	return db
}

func NewDriver(dbPath string) IDriverSqlite {
	return &driver{
		File: dbPath,
	}
}
