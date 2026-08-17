package users

import (
	"net/http"

	"github.com/YomikaiHub/ohara/api/internal/http/auth"
	"github.com/YomikaiHub/ohara/api/internal/httputil"
	"github.com/YomikaiHub/ohara/api/internal/utils"
)

func (h *Handler) UserInfo(w http.ResponseWriter, req *http.Request) {
	userID, ok := req.Context().Value(utils.UserIDKey).(string)
	if !ok || userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.store.User.GetUserByID(req.Context(), userID)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
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
		httputil.InternalServerError(h.logger, w, req, err)
	}
}
