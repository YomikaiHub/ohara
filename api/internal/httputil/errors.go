package httputil

import (
	"log/slog"
	"net/http"
)

func InternalServerError(
	logger *slog.Logger,
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	logger.Error(
		"internal server error",
		slog.Int("status", http.StatusInternalServerError),
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
		slog.String("ip", r.RemoteAddr),
		slog.String("error", err.Error()),
	)

	_ = WriteErrorJSON(
		w,
		http.StatusInternalServerError,
		"INTERNAL_SERVER_ERROR",
		"The server encountered an internal error.",
	)
}

func BadRequest(
	logger *slog.Logger,
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	message := "Bad request."

	if err != nil {
		message = err.Error()
	}

	logger.Warn(
		"bad request",
		slog.Int("status", http.StatusBadRequest),
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
		slog.String("ip", r.RemoteAddr),
		slog.String("error", message),
	)

	_ = WriteErrorJSON(
		w,
		http.StatusBadRequest,
		"BAD_REQUEST",
		message,
	)
}

func Unauthorized(
	logger *slog.Logger,
	w http.ResponseWriter,
	r *http.Request,
	message string,
) {
	if message == "" {
		message = "Invalid or expired authentication token."
	}

	logger.Warn(
		"unauthorized access",
		slog.Int("status", http.StatusUnauthorized),
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
		slog.String("ip", r.RemoteAddr),
		slog.String("error", message),
	)

	_ = WriteErrorJSON(
		w,
		http.StatusUnauthorized,
		"UNAUTHORIZED",
		message,
	)
}

func NotFound(
	logger *slog.Logger,
	w http.ResponseWriter,
	r *http.Request,
) {
	logger.Warn(
		"resource not found",
		slog.Int("status", http.StatusNotFound),
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
		slog.String("ip", r.RemoteAddr),
	)

	_ = WriteErrorJSON(
		w,
		http.StatusNotFound,
		"RESOURCE_NOT_FOUND",
		"The requested resource was not found.",
	)
}
