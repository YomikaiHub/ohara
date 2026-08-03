package store

import (
	"context"
	"database/sql"

	"github.com/YomikaiHub/ohara/api/internal/errs"
	"github.com/YomikaiHub/ohara/api/internal/models"
)

type SessionStore struct {
	db *sql.DB
}

func (store *SessionStore) Create(
	ctx context.Context,
	session *models.Session,
) error {
	query := `
		INSERT INTO sessions (user_id, refresh_token_hash, expires_at) VALUES ($1, $2, $3)
		RETURNING id, user_id, refresh_token_hash, expires_at, revoked_at, created_at, updated_at;
	`

	err := store.db.QueryRowContext(
		ctx,
		query,
		session.UserID,
		session.RefreshTokenHash,
		session.ExpiresAt,
	).Scan(
		&session.ID,
		&session.UserID,
		&session.RefreshTokenHash,
		&session.ExpiresAt,
		&session.RevokedAt,
		&session.CreatedAt,
		&session.UpdatedAt,
	)
	if err != nil {
		return errs.HandleSessionError(err)
	}

	return nil
}
