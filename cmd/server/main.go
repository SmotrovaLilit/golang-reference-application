package main

import (
	"fmt"
	"os"
)

func main() {
	app := NewApplication()
	err := app.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
