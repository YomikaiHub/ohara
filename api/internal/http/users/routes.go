package users

import (
	"log/slog"

	"github.com/YomikaiHub/ohara/api/internal/config"
	"github.com/YomikaiHub/ohara/api/internal/middleware"
	"github.com/YomikaiHub/ohara/api/internal/store"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(
	r chi.Router,
	cfg *config.Config,
	logger *slog.Logger,
	store *store.Store,
) {
	h := New(cfg, logger, store)

	r.Route("/users", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware([]byte(h.config.JWT.Key)))

		r.Get("/me", h.UserInfo)
		r.Patch("/me", h.UpdateUser)
	})
}
