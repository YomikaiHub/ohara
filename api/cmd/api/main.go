package main

import (
	stdlog "log"

	"github.com/YomikaiHub/ohara/api/internal/db"
	"github.com/YomikaiHub/ohara/api/internal/env"
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

	cfg := &config{
		addr:    env.GetString("ADDR", ":8080"),
		env:     env.MustGet("ENV"),
		version: env.GetString("VERSION", "0.0.1"),
		db: dbConfig{
			dbUrl:          env.MustGet("DB_URL"),
			dbMaxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			dbMaxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			dbMaxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15min"),
		},
	}

	log := logger.New(cfg.env)

	db, err := db.New(
		cfg.db.dbUrl,
		cfg.db.dbMaxOpenConns,
		cfg.db.dbMaxIdleConns,
		cfg.db.dbMaxIdleTime,
	)
	if err != nil {
		stdlog.Panic(err)
	}
	defer db.Close()

	store := store.NewDBStorage(db)

	app := application{
		config: *cfg,
		logger: log,
		store:  store,
	}

	app.logger.Info("application starting....")

	mux := app.mount()

	err = app.run(mux)
	if err != nil {
		app.logger.Error(err.Error())
	}
}
