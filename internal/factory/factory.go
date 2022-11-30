package factory

import (
	"log"

	"github.com/hrz8/gokomodo-challenge/internal/driver/http"
	"github.com/hrz8/gokomodo-challenge/internal/driver/sqlite"
)

type (
	factory struct {
		Driver struct {
			Http   http.IDriverHttp
			Sqlite sqlite.IDriverSqlite
		}
	}

	IFactory interface {
		Start()
	}
)

func (f *factory) Start() {
	conn := f.Driver.Sqlite.Start()
	err := f.Driver.Http.Start(conn)
	if err != nil {
		log.Fatal("Failed to start the app!")
	}
}

func NewFactory(
	driverHttp http.IDriverHttp,
	driverSqlite sqlite.IDriverSqlite,
) IFactory {
	factory := &factory{}

	factory.Driver.Http = driverHttp
	factory.Driver.Sqlite = driverSqlite

	return factory
}
