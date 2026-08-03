package api

import (
	"net/http"
	"time"

	"github.com/YomikaiHub/ohara/api/internal/http/auth"
	"github.com/YomikaiHub/ohara/api/internal/http/health"
	"github.com/YomikaiHub/ohara/api/internal/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(logger.Middleware(app.logger))
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Second * 60))

	r.Route("/v1", func(r chi.Router) {
		health.RegisterRoutes(r, app.config, app.logger)
		auth.RegisterRoutes(r, app.config, app.logger, app.store)
	})

	return r
}
