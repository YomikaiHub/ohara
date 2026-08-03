package auth

import (
	"log/slog"

	"github.com/YomikaiHub/ohara/api/internal/config"
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

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register/email", h.Register)
		r.Post("/login/email", h.LoginWithEmail)
	})
}
