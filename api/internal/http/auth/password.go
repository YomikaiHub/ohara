package auth

import (
	"net/http"

	"github.com/YomikaiHub/ohara/api/internal/httputil"
	"github.com/YomikaiHub/ohara/api/internal/utils"
	"github.com/alexedwards/argon2id"
)

func (h *Handler) ChangePassword(
	w http.ResponseWriter,
	req *http.Request,
) {
	var payload ChangePasswordRequest

	if err := httputil.ReadJSON(w, req, &payload); err != nil {
		httputil.BadRequest(h.logger, w, req, err)
		return
	}

	userID, ok := req.Context().Value(utils.UserIDKey).(string)
	if !ok || userID == "" {
		httputil.Unauthorized(h.logger, w, req, "Invalid or expired authentication token")
		return
	}

	account, err := h.store.Auth.GetCredentialAccountByUserID(
		req.Context(),
		userID,
	)
	if err != nil {
		httputil.InternalServerError(h.logger, w, req, err)
		return
	}

	isPasswordCorrect, err := VerifyPassword(
		payload.CurrentPassword,
		account.PasswordHash,
	)
	if err != nil {
		httputil.InternalServerError(h.logger, w, req, err)
		return
	}

	if !isPasswordCorrect {
		httputil.Unauthorized(h.logger, w, req, "Invalid email or password.")
		return
	}

	newPasswordHash, err := HashPassword(payload.NewPassword)
	if err != nil {
		httputil.InternalServerError(h.logger, w, req, err)
		return
	}

	if err := h.store.Auth.ChangePassword(
		req.Context(),
		userID,
		newPasswordHash,
	); err != nil {

		httputil.InternalServerError(h.logger, w, req, err)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func HashPassword(password string) (string, error) {
	params := &argon2id.Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	hash, err := argon2id.CreateHash(password, params)
	if err != nil {
		return "", err
	}

	return hash, nil
}

func VerifyPassword(password, hash string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false, err
	}

	return match, nil
}
