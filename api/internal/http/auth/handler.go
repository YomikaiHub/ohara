package auth

import (
	"log/slog"

	"github.com/YomikaiHub/ohara/api/internal/config"
	"github.com/YomikaiHub/ohara/api/internal/store"
)

type Handler struct {
	config *config.Config
	logger *slog.Logger
	store  *store.Store
}

func New(
	cfg *config.Config,
	logger *slog.Logger,
	store *store.Store,
) *Handler {
	return &Handler{
		config: cfg,
		logger: logger,
		store:  store,
	}
}
