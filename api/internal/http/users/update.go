package users

import (
	"errors"
	"net/http"

	"github.com/YomikaiHub/ohara/api/internal/errs"
	"github.com/YomikaiHub/ohara/api/internal/http/auth"
	"github.com/YomikaiHub/ohara/api/internal/httputil"
	"github.com/YomikaiHub/ohara/api/internal/utils"
)

func (h *Handler) UpdateUser(w http.ResponseWriter, req *http.Request) {
	var payload UpdateUserRequest

	if err := httputil.ReadJSON(w, req, &payload); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	userID, ok := req.Context().Value(utils.UserIDKey).(string)
	if !ok || userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.store.User.Update(
		req.Context(),
		userID,
		payload.Username,
		payload.FirstName,
		payload.LastName,
		payload.Image,
	)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrUserNotFound):
			httputil.NotFound(h.logger, w, req)

		case errors.Is(err, errs.ErrUsernameAlreadyExists):
			httputil.BadRequest(h.logger, w, req, err)

		case errors.Is(err, errs.ErrNoFieldsToUpdate):
			httputil.BadRequest(h.logger, w, req, err)

		default:
			httputil.InternalServerError(h.logger, w, req, err)
		}
		return
	}

	response := auth.UserResponsePayload{
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

	if err := httputil.WriteJSON(w, http.StatusOK, response); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
