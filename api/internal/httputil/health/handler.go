package health

import (
	"log/slog"

	"github.com/YomikaiHub/ohara/api/internal/config"
)

type Handler struct {
	config *config.Config
	logger *slog.Logger
}

func New(
	cfg *config.Config,
	logger *slog.Logger,
) *Handler {
	return &Handler{
		config: cfg,
		logger: logger,
	}
}
