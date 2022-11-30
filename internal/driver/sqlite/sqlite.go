package sqlite

import (
	"log"

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

	return db
}

func NewDriver(dbPath string) IDriverSqlite {
	return &driver{
		File: dbPath,
	}
}
