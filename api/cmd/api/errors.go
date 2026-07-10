package main

import (
	"log/slog"
	"net/http"
)

func (app *application) internalServerError(
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	app.logger.Error(
		"internal server error",
		slog.Int("status", http.StatusInternalServerError),
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
		slog.String("ip", r.RemoteAddr),
		slog.String("error", err.Error()),
	)

	_ = writeErrorJSON(
		w,
		http.StatusInternalServerError,
		"INTERNAL_SERVER_ERROR",
		"The server encountered an internal error.",
	)
}

func (app *application) badRequestError(
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	message := "Bad request."

	if err != nil {
		message = err.Error()
	}

	app.logger.Warn(
		"bad request",
		slog.Int("status", http.StatusBadRequest),
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
		slog.String("ip", r.RemoteAddr),
		slog.String("error", message),
	)

	_ = writeErrorJSON(
		w,
		http.StatusBadRequest,
		"BAD_REQUEST",
		message,
	)
}

func (app *application) notFoundError(
	w http.ResponseWriter,
	r *http.Request,
) {
	app.logger.Warn(
		"resource not found",
		slog.Int("status", http.StatusNotFound),
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
		slog.String("ip", r.RemoteAddr),
	)

	_ = writeErrorJSON(
		w,
		http.StatusNotFound,
		"RESOURCE_NOT_FOUND",
		"The requested resource was not found.",
	)
}

func (app *application) unauthorizedError(
	w http.ResponseWriter,
	r *http.Request,
) {
	app.logger.Warn(
		"unauthorized access",
		slog.Int("status", http.StatusUnauthorized),
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
		slog.String("ip", r.RemoteAddr),
	)

	_ = writeErrorJSON(
		w,
		http.StatusUnauthorized,
		"UNAUTHORIZED",
		"Invalid or expired authentication token.",
	)
}
