package main

import (
	"flag"
	"fmt"
	"os"
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

	app := NewApplication(
		ConnectDB(dbType, dsn),
		HTTPAddr(httpAddr),
	)
	err := app.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
