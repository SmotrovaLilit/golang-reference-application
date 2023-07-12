//go:generate go run github.com/google/wire/cmd/wire@latest
//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"reference-application/internal/application"
	"reference-application/internal/application/commands/createprogram"
	"reference-application/internal/infrastructure/repositories"
	infrastructurehttp "reference-application/internal/infrastructure/transports/http"
)

type Application struct {
	HTTPHandler http.Handler
}

func (app Application) Run() error {
	return http.ListenAndServe(":8080", app.HTTPHandler)
}

func NewApplication() Application {
	wire.Build(
		wire.Struct(new(Application), "*"),
		infrastructurehttp.NewHandler,
		application.Set,
		createprogram.Set,
		repositories.Set,
		provideDB,
	)
	return Application{}
}

func provideDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(repositories.ProgramModel{})
	if err != nil {
		panic("failed to migrate database")
	}
	return db
}
