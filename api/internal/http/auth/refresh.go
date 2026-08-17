package auth

import (
	"errors"
	"net/http"

	"github.com/YomikaiHub/ohara/api/internal/errs"
	"github.com/YomikaiHub/ohara/api/internal/httputil"
	"github.com/YomikaiHub/ohara/api/internal/utils"
)

func (h *Handler) RefreshToken(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("yomikai_refresh")
	if err != nil {
		httputil.Unauthorized(h.logger, w, req, "")
		return
	}

	refreshTokenHash := utils.HashRefreshToken(cookie.Value)

	session, user, err := h.store.Session.GetSessionWithUserByRefreshTokenHash(req.Context(), refreshTokenHash)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrSessionNotFound):
			httputil.Unauthorized(
				h.logger,
				w,
				req,
				"",
			)
		default:
			httputil.InternalServerError(
				h.logger,
				w,
				req,
				err,
			)
		}
		return
	}

	if err := h.ValidateSession(session); err != nil {
		httputil.Unauthorized(h.logger, w, req, "")
		return
	}

	if err := h.RotateSession(w, req, user, session); err != nil {
		httputil.InternalServerError(h.logger, w, req, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
