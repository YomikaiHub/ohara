package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/YomikaiHub/ohara/api/internal/logger"
	"github.com/YomikaiHub/ohara/api/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	logger *slog.Logger
	store  store.Storage
}

type config struct {
	addr    string
	db      dbConfig
	env     string
	version string
}

type dbConfig struct {
	dbUrl          string
	dbMaxOpenConns int
	dbMaxIdleConns int
	dbMaxIdleTime  string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(logger.Middleware(app.logger))
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Second * 60))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	app.logger.Info("server running", "addr", app.config.addr)

	return srv.ListenAndServe()
}
