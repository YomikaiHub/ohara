package errs

import (
	"errors"

	"github.com/lib/pq"
)

func HandleUserError(err error) error {
	var pqErr *pq.Error

	if errors.As(err, &pqErr) && pqErr.Code == "23505" {
		switch pqErr.Constraint {
		case "users_email_key":
			return ErrEmailAlreadyExists
		case "users_username_key":
			return ErrUsernameAlreadyExists
		}
	}

	return err
}

func HandleAccountError(err error) error {
	var pqErr *pq.Error

	if errors.As(err, &pqErr) && pqErr.Code == "23505" {
		switch pqErr.Constraint {
		case "accounts_provider_id_account_id_key":
			return ErrAccountAlreadyExists
		}
	}

	return err
}

func HandleSessionError(err error) error {
	var pqErr *pq.Error

	if errors.As(err, &pqErr) {
		switch pqErr.Code {
		case "23505":
			if pqErr.Constraint == "sessions_refresh_token_hash_key" {
				return ErrSessionAlreadyExists
			}

		case "23503":
			return ErrUserNotFound
		}
	}

	return err
}

func HandleUserUpdateError(err error) error {
	var pqErr *pq.Error

	if errors.As(err, &pqErr) {
		switch pqErr.Code {
		case "23505":
			if pqErr.Constraint == "users_username_key" {
				return ErrUsernameAlreadyExists
			}

		case "23503":
			return ErrUserNotFound
		}
	}

	return err
}
