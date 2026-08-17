package auth

import (
	"net/http"

	"github.com/YomikaiHub/ohara/api/internal/httputil"
	"github.com/YomikaiHub/ohara/api/internal/utils"
)

func (h *Handler) Logout(w http.ResponseWriter, req *http.Request) {
	sessionID, ok := req.Context().
		Value(utils.SessionIDKey).(string)

	if !ok || sessionID == "" {
		httputil.Unauthorized(
			h.logger,
			w,
			req,
			"Invalid or expired authentication token.",
		)
		return
	}

	if err := h.store.Session.Revoke(
		req.Context(),
		sessionID,
	); err != nil {
		httputil.InternalServerError(
			h.logger,
			w,
			req,
			err,
		)
		return
	}

	h.clearAuthCookies(w)

	w.WriteHeader(http.StatusNoContent)
}
