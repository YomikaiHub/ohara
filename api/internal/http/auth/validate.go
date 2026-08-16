package auth

import (
	"time"

	"github.com/YomikaiHub/ohara/api/internal/errs"
	"github.com/YomikaiHub/ohara/api/internal/models"
)

func (h *Handler) ValidateSession(session *models.Session) error {
	if session == nil {
		return errs.ErrSessionNotFound
	}

	if session.RevokedAt != nil {
		return errs.ErrSessionRevoked
	}

	if time.Now().After(session.ExpiresAt) {
		return errs.ErrSessionExpired
	}

	return nil
}
