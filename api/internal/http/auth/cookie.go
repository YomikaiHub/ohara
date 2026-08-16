package auth

import (
	"net/http"
	"strings"
	"time"

	"github.com/YomikaiHub/ohara/api/internal/models"
)

func (h *Handler) CreateSession(
	w http.ResponseWriter,
	req *http.Request,
	user *models.User,
) error {
	refreshToken, err := GenerateRefreshToken(h.config.JWT.RefreshTokenTTL)
	if err != nil {
		return err
	}

	session := &models.Session{
		UserID:           user.ID,
		RefreshTokenHash: refreshToken.Hash,
		ExpiresAt:        refreshToken.ExpiresAt,
	}

	err = h.store.Session.Create(req.Context(), session)
	if err != nil {
		return err
	}

	accessToken, err := GenerateAccessToken(
		user.ID,
		session.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		[]byte(h.config.JWT.Key),
		h.config.JWT.AccessTokenTTL,
	)
	if err != nil {
		return err
	}

	h.setCookies(
		w,
		accessToken,
		refreshToken.PlainText,
		session.ExpiresAt,
	)

	return nil
}

func (h *Handler) RotateSession(
	w http.ResponseWriter,
	req *http.Request,
	user *models.User,
	session *models.Session,
) error {
	refreshToken, err := GenerateRefreshToken(h.config.JWT.RefreshTokenTTL)
	if err != nil {
		return err
	}

	accessToken, err := GenerateAccessToken(
		user.ID,
		session.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		[]byte(h.config.JWT.Key),
		h.config.JWT.AccessTokenTTL,
	)
	if err != nil {
		return err
	}

	if err := h.store.Session.RotateRefreshToken(
		req.Context(),
		session.ID,
		refreshToken.Hash,
		refreshToken.ExpiresAt,
	); err != nil {
		return err
	}

	session.RefreshTokenHash = refreshToken.Hash
	session.ExpiresAt = refreshToken.ExpiresAt

	h.setCookies(
		w,
		accessToken,
		refreshToken.PlainText,
		refreshToken.ExpiresAt,
	)

	return nil
}

func (h *Handler) setCookies(
	w http.ResponseWriter,
	accessToken string,
	refreshToken string,
	refreshExpiry time.Time,
) {
	isCookieSecure := strings.ToLower(h.config.App.Env) == "production"
	now := time.Now()

	http.SetCookie(
		w,
		&http.Cookie{
			Name:     "yomikai_access",
			Value:    accessToken,
			Path:     "/",
			HttpOnly: true,
			Secure:   isCookieSecure,
			SameSite: http.SameSiteLaxMode,
			MaxAge:   int(h.config.JWT.AccessTokenTTL.Seconds()),
			Expires:  now.Add(h.config.JWT.AccessTokenTTL),
		},
	)

	http.SetCookie(
		w,
		&http.Cookie{
			Name:     "yomikai_refresh",
			Value:    refreshToken,
			Path:     "/api/v1/auth/refresh",
			HttpOnly: true,
			Secure:   isCookieSecure,
			SameSite: http.SameSiteLaxMode,
			MaxAge:   int(h.config.JWT.RefreshTokenTTL.Seconds()),
			Expires:  refreshExpiry,
		},
	)
}
