package main

import (
	"cash-flow/server/internal/bootstrap"
)

func main() {
	app := bootstrap.NewApp()
	app.Run()
}
