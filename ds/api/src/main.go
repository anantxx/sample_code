package main

import (
	"github.com/anant/final/ds/api/app"
	"github.com/anant/final/ds/api/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3030")
}
