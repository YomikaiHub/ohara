package logger

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

func New(env string) *slog.Logger {
	file, err := openWeeklyLogFile()
	if err != nil {
		panic(err)
	}

	level := slog.LevelDebug
	if env == "production" {
		level = slog.LevelInfo
	}

	if env == "production" {
		return slog.New(
			&MultiHandler{
				handlers: []slog.Handler{
					slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
						Level:       level,
						AddSource:   true,
						ReplaceAttr: replaceSource,
					}),
					slog.NewJSONHandler(file, &slog.HandlerOptions{
						Level:       level,
						AddSource:   true,
						ReplaceAttr: replaceSource,
					}),
				},
			},
		)
	}
	return slog.New(
		&MultiHandler{
			handlers: []slog.Handler{
				tint.NewHandler(os.Stdout, &tint.Options{
					Level:       level,
					TimeFormat:  "15:04:05",
					AddSource:   true,
					ReplaceAttr: replaceSource,
				}),
				slog.NewTextHandler(file, &slog.HandlerOptions{
					Level:       level,
					AddSource:   true,
					ReplaceAttr: replaceSource,
				}),
			},
		},
	)
}
