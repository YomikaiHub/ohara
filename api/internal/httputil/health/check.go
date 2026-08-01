package health

import (
	"net/http"

	"github.com/YomikaiHub/ohara/api/internal/httputil"
)

func (h *Handler) Check(w http.ResponseWriter, req *http.Request) {
	response := Response{
		Status:      "available",
		Environment: h.config.App.Env,
		Version:     h.config.App.Version,
	}

	if err := httputil.WriteJSON(w, http.StatusOK, response); err != nil {
		h.logger.Error("failed to write health response", "error", err)

		_ = httputil.WriteErrorJSON(
			w,
			http.StatusInternalServerError,
			"INTERNAL_SERVER_ERROR",
			"An unexpected error occurred.",
		)
	}
}
