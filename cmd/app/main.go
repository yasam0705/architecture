package main

import (
	"github/architecture/config"
	"github/architecture/internal/app"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	err = app.Run(cfg)
	if err != nil {
		panic(err)
	}

	log.Println(cfg.App, "app service stops")
}
