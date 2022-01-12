package main

import (
	"context"
	"log"

	"com.thebeachmaster/entexample/config"
	"com.thebeachmaster/entexample/ent"
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

	// postgresql://user:password@host/database
	// TODO: Handle Heroku DB Conn String
	dbConnString := "postgresql://" + cfg.DB.DBUser + ":" + cfg.DB.DBPass + "@" + cfg.DB.DBHost + "/" + cfg.DB.DBName

	client, err := ent.Open("postgres", dbConnString)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Add ent Client
	server := server.NewServer(cfg, client)
	if err = server.Run(); err != nil {
		log.Fatal(err)
	}
}
