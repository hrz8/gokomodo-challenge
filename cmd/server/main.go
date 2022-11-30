package main

import (
	"github.com/hrz8/gokomodo-challenge/internal/driver/http"
	"github.com/hrz8/gokomodo-challenge/internal/driver/sqlite"
	"github.com/hrz8/gokomodo-challenge/internal/factory"
)

func main() {
	httpDriver := http.NewDriver()
	sqliteDriver := sqlite.NewDriver("db.sqlite")

	factory := factory.NewFactory(httpDriver, sqliteDriver)

	factory.Start()
}
