//go:generate go run github.com/google/wire/cmd/wire@latest
//go:build wireinject
// +build wireinject

package main

import (
	"fmt"
	"github.com/google/wire"
	"gorm.io/gorm"
	"net"
	"net/http"
	"reference-application/internal/application"
	"reference-application/internal/application/commands/createprogram"
	"reference-application/internal/application/commands/sendtoreviewprogramversion"
	"reference-application/internal/application/commands/updateprogramversion"
	"reference-application/internal/infrastructure/repositories"
	infrastructurehttp "reference-application/internal/infrastructure/transports/http"
)

type HTTPAddr string

type Application struct {
	HTTPHandler http.Handler
	HTTPAddr    HTTPAddr
}

func (app *Application) Run() error {
	ln, err := net.Listen("tcp", string(app.HTTPAddr))
	if err != nil {
		return err
	}
	fmt.Println("HTTP server listening on", string(app.HTTPAddr))
	return app.Serve(ln)
}

func (app *Application) Serve(l net.Listener) error {
	return http.Serve(l, app.HTTPHandler)
}

func NewApplication(
	db *gorm.DB,
	addr HTTPAddr,
) Application {
	wire.Build(
		wire.Struct(new(Application), "*"),
		infrastructurehttp.NewHandler,
		application.Set,
		createprogram.Set,
		updateprogramversion.Set,
		sendtoreviewprogramversion.Set,
		repositories.Set,
	)
	return Application{}
}
