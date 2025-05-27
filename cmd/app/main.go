package main

import (
	"log"

	"github.com/zinct/amanmemilih/config"
	"github.com/zinct/amanmemilih/internal/app"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
