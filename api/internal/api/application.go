package api

import (
	"log/slog"

	"github.com/YomikaiHub/ohara/api/internal/config"
	"github.com/YomikaiHub/ohara/api/internal/store"
)

type Application struct {
	config *config.Config
	logger *slog.Logger
	store  *store.Store
}

func New(
	cfg *config.Config,
	log *slog.Logger,
	store *store.Store,
) *Application {
	return &Application{
		config: cfg,
		logger: log,
		store:  store,
	}
}
