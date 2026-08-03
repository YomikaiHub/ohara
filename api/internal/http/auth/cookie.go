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
	accessToken, err := GenerateAccessToken(
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		[]byte(h.config.JWT.Key),
		h.config.JWT.AccessTokenTTL,
	)
	if err != nil {
		return err
	}

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

	http.SetCookie(w, &http.Cookie{
		Name:     "yomikai_refresh",
		Value:    refreshToken.PlainText,
		HttpOnly: true,
		Secure:   isCookieSecure,
		SameSite: http.SameSiteLaxMode,
		Path:     "/api/v1/auth/refresh",
		MaxAge:   int(h.config.JWT.RefreshTokenTTL.Seconds()),
		Expires:  now.Add(h.config.JWT.RefreshTokenTTL),
	})

	return nil
}
