package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"com.thebeachmaster/entexample/config"
	"com.thebeachmaster/entexample/ent"
	"com.thebeachmaster/entexample/server"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func openDb(databaseUrl string) (*ent.Client, error) {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		return nil, err
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), nil
}

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
	connStringBuilder := strings.Builder{}
	connStringBuilder.WriteString(fmt.Sprintf("postgresql://%s:%s@%s/%s", cfg.DB.DBUser, cfg.DB.DBPass, cfg.DB.DBHost, cfg.DB.DBName))
	dbConnString := connStringBuilder.String()
	client, err := openDb(dbConnString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	// TODO: maybe add cli flag
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	server := server.NewServer(cfg, client)
	if err = server.Run(); err != nil {
		log.Fatal(err)
	}
}
