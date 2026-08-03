package auth

import (
	"net/http"

	"github.com/YomikaiHub/ohara/api/internal/httputil"
)

func (h *Handler) LoginWithEmail(w http.ResponseWriter, req *http.Request) {
	var payload LoginRequest

	if err := httputil.ReadJSON(w, req, &payload); err != nil {
		httputil.BadRequest(h.logger, w, req, err)
		return
	}

	user, account, err := h.store.Auth.GetUserWithCredentialByEmail(req.Context(), payload.Email)
	if err != nil {
		httputil.Unauthorized(
			h.logger,
			w,
			req,
			"Invalid email or password.",
		)

		return
	}

	isPasswordCorrect, err := VerifyPassword(payload.Password, account.PasswordHash)
	if err != nil {
		httputil.InternalServerError(h.logger, w, req, err)
		return
	}

	if !isPasswordCorrect {
		httputil.Unauthorized(
			h.logger,
			w,
			req,
			"Invalid email or password.",
		)

		return
	}

	if err := h.CreateSession(w, req, user); err != nil {
		httputil.InternalServerError(h.logger, w, req, err)
		return
	}

	response := UserResponsePayload{
		ID:            user.ID,
		Email:         user.Email,
		Username:      user.Username,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Image:         user.Image,
		EmailVerified: user.EmailVerified,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}

	if err := httputil.WriteJSON(
		w,
		http.StatusOK,
		response,
	); err != nil {
		httputil.InternalServerError(h.logger, w, req, err)
		return
	}
}
