package api

import (
	"net/http"
	"time"
)

func (app *Application) Run() error {
	mux := app.mount()

	srv := &http.Server{
		Addr:         app.config.App.Addr,
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	app.logger.Info("application starting")
	app.logger.Info("server running", "addr", app.config.App.Addr)

	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
