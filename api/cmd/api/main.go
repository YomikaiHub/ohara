package main

import (
	stdlog "log"
	"os"

	"github.com/YomikaiHub/ohara/api/internal/api"
	"github.com/YomikaiHub/ohara/api/internal/config"
	"github.com/YomikaiHub/ohara/api/internal/db"
	"github.com/YomikaiHub/ohara/api/internal/logger"
	"github.com/YomikaiHub/ohara/api/internal/store"
	"github.com/joho/godotenv"
)

const version = "0.0.1"

func main() {
	err := godotenv.Load()
	if err != nil {
		stdlog.Fatal("No env vars file found!")
	}

	cfg := config.Load(version)

	log := logger.New(cfg.App.Env)

	dbConn, err := db.New(cfg.DB)
	if err != nil {
		log.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	log.Info("database connected")
	defer dbConn.Close()

	store := store.New(dbConn)

	app := api.New(cfg, log, store)
	if err := app.Run(); err != nil {
		log.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
