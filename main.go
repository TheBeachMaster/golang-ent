package main

import (
	"log"

	"com.thebeachmaster/entexample/config"
	"com.thebeachmaster/entexample/server"
)

func main() {
	log.Println("Starting Server...")

	appConfigPath := "./config/config-local"

	cfgFile, err := config.LoadConfig(appConfigPath)
	if err != nil {
		log.Fatalf("LoadConfig Error: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig Error: %v", err)
	}

	server := server.NewServer(cfg)
	if err = server.Run(); err != nil {
		log.Fatal(err)
	}
}
