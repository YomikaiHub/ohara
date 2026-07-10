package logger

import (
	"context"
	"log/slog"
)

type MultiHandler struct {
	handlers []slog.Handler
}

func (mh *MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, handler := range mh.handlers {
		if handler.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

func (mh *MultiHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, handler := range mh.handlers {
		if err := handler.Handle(ctx, r); err != nil {
			return err
		}
	}
	return nil
}

func (mh *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	handlers := make([]slog.Handler, len(mh.handlers))

	for i, handler := range mh.handlers {
		handlers[i] = handler.WithAttrs(attrs)
	}

	return &MultiHandler{
		handlers: handlers,
	}
}

func (mh *MultiHandler) WithGroup(name string) slog.Handler {
	handlers := make([]slog.Handler, len(mh.handlers))

	for i, handler := range mh.handlers {
		handlers[i] = handler.WithGroup(name)
	}

	return &MultiHandler{
		handlers: handlers,
	}
}
