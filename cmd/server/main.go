package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"reference-application/internal/pkg/log"
)

var httpAddr = "localhost:8080"
var dsn = "test.db"
var dbType = "sqlite"

var flags = flag.NewFlagSet("server", flag.ExitOnError)

func main() {
	flags.StringVar(&httpAddr, "addr", httpAddr, "HTTP server address")
	flags.StringVar(&dsn, "dsn", dsn, "Database DSN")
	flags.StringVar(&dbType, "db", dbType, "Database type")
	_ = flags.Parse(os.Args[1:])
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug, // get from config TODO https://github.com/SmotrovaLilit/golang-reference-application/issues/31
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	logger = log.WithApplicationInfo(logger, os.Getenv("APPLICATION_NAME"))
	db, err := ConnectDB(dbType, dsn)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to connect to database: %s", err.Error()))
		os.Exit(1)
	}
	app, err := NewApplication(
		db,
		HTTPAddr(httpAddr),
		logger,
	)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to create application: %s", err.Error()))
		os.Exit(1)
	}
	err = app.Run()
	if err != nil {
		logger.Error(fmt.Sprintf("failed to run application: %s", err))
		os.Exit(1)
	}
}
