package auth

import (
	"errors"
	"net/http"

	"github.com/YomikaiHub/ohara/api/internal/errs"
	"github.com/YomikaiHub/ohara/api/internal/httputil"
	"github.com/YomikaiHub/ohara/api/internal/models"
)

func (h *Handler) Register(w http.ResponseWriter, req *http.Request) {
	var payload RegisterRequest

	if err := httputil.ReadJSON(w, req, &payload); err != nil {
		httputil.BadRequest(h.logger, w, req, err)
		return
	}

	hashedPassword, err := HashPassword(payload.Password)
	if err != nil {
		httputil.InternalServerError(h.logger, w, req, err)
		return
	}

	user := &models.User{
		Email:     payload.Email,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Username:  payload.Username,
		Image:     payload.Image,
	}

	account := &models.Account{
		AccountID:    payload.Email,
		ProviderID:   "credential",
		PasswordHash: hashedPassword,
	}

	err = h.store.Auth.RegisterWithEmail(req.Context(), user, account)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrEmailAlreadyExists):
			httputil.BadRequest(h.logger, w, req, errs.ErrEmailAlreadyExists)
		case errors.Is(err, errs.ErrAccountAlreadyExists):
			httputil.BadRequest(h.logger, w, req, errs.ErrAccountAlreadyExists)
		case errors.Is(err, errs.ErrUsernameAlreadyExists):
			httputil.BadRequest(h.logger, w, req, errs.ErrUsernameAlreadyExists)
		default:
			httputil.InternalServerError(h.logger, w, req, err)
		}

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
		http.StatusCreated,
		response,
	); err != nil {
		httputil.InternalServerError(h.logger, w, req, err)
		return
	}
}
