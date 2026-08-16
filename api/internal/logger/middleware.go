package logger

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type contextKey string

const RequestIDKey contextKey = "request_id"

func Middleware(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := uuid.NewString()
			start := time.Now()

			ctx := context.WithValue(
				r.Context(),
				RequestIDKey,
				requestID,
			)

			w.Header().Set("X-Request-ID", requestID)

			rw := &responseWriter{
				ResponseWriter: w,
				status:         http.StatusOK,
			}

			next.ServeHTTP(rw, r.WithContext(ctx))

			log.Info(
				"request completed",
				slog.String("request_id", requestID),
				slog.Int("status", rw.status),
				slog.String("method", r.Method),
				slog.String("url", r.URL.Path),
				slog.String("user_agent", r.UserAgent()),
				slog.Duration("duration", time.Since(start)),
			)
		})
	}
}
