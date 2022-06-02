package main

import (
	"log"
	"majezanu/capstone/config"
	"majezanu/capstone/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config custom_error: %s", err)
	}

	// Run
	app.Run(cfg)
}
