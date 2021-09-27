package main

import (
	"github.com/cubetiq/zengo/app"
	"github.com/cubetiq/zengo/config"
)

func main() {
	config := config.GetConfig()

	app := app.App{}
	app.Run(config)
}
