package sqlite

import (
	"fmt"
	"log"
	"os"

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
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	db, err := gorm.Open(sqliteDriver.Open(d.File), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to open connection to database!")
	}

	err = db.AutoMigrate(&entity.Buyer{})
	if err != nil {
		log.Fatal("Failed run migration!")
	}

	err = db.AutoMigrate(&entity.Seller{})
	if err != nil {
		log.Fatal("Failed run migration!")
	}

	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		log.Fatal("Failed run migration!")
	}

	err = db.AutoMigrate(&entity.Order{})
	if err != nil {
		log.Fatal("Failed run migration!")
	}

	err = db.AutoMigrate(&entity.OrderItem{})
	if err != nil {
		log.Fatal("Failed run migration!")
	}

	return db
}

func NewDriver(dbPath string) IDriverSqlite {
	return &driver{dbPath}
}
