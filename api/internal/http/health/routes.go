package health

import (
	"log/slog"

	"github.com/YomikaiHub/ohara/api/internal/config"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(
	r chi.Router,
	cfg *config.Config,
	logger *slog.Logger,
) {
	h := New(cfg, logger)

	r.Get("/health", h.Check)
}
