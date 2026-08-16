package store

import (
	"context"
	"database/sql"
	"time"

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

func (store *SessionStore) GetByRefreshTokenHash(
	ctx context.Context,
	hash string,
) (*models.Session, error) {
	query := `
		SELECT id, user_id, refresh_token_hash, expires_at, revoked_at, created_at, updated_at
		FROM sessions
		WHERE refresh_token_hash = $1;
	`

	session := &models.Session{}

	err := store.db.QueryRowContext(
		ctx,
		query,
		hash,
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
		return nil, errs.HandleSessionError(err)
	}

	return session, nil
}

func (store *SessionStore) RotateRefreshToken(
	ctx context.Context,
	sessionID string,
	newHash string,
	expiresAt time.Time,
) error {
	query := `
		UPDATE sessions
		SET
    		refresh_token_hash = $1,
    		expires_at = $2,
    		updated_at = NOW()
		WHERE id = $3;
	`

	result, err := store.db.ExecContext(
		ctx,
		query,
		newHash,
		expiresAt,
		sessionID,
	)
	if err != nil {
		return errs.HandleSessionError(err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (store *SessionStore) Revoke(
	ctx context.Context, sessionID string,
) error {
	query := `
		DELETE FROM sessions WHERE id = $1
	`

	result, err := store.db.ExecContext(
		ctx,
		query,
		sessionID,
	)
	if err != nil {
		return errs.HandleSessionError(err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (store *SessionStore) GetSessionWithUserByRefreshTokenHash(
	ctx context.Context,
	hash string,
) (*models.Session, *models.User, error) {
	query := `
		SELECT
			s.id,
			s.user_id,
			s.refresh_token_hash,
			s.expires_at,
			s.revoked_at,
			s.created_at,
			s.updated_at,

			u.id,
			u.email,
			u.username,
			u.first_name,
			u.last_name,
			u.image,
			u.email_verified,
			u.created_at,
			u.updated_at
		FROM sessions s
		JOIN users u
			ON s.user_id = u.id
		WHERE s.refresh_token_hash = $1;
	`

	session := &models.Session{}
	user := &models.User{}

	err := store.db.QueryRowContext(
		ctx,
		query,
		hash,
	).Scan(
		&session.ID,
		&session.UserID,
		&session.RefreshTokenHash,
		&session.ExpiresAt,
		&session.RevokedAt,
		&session.CreatedAt,
		&session.UpdatedAt,

		&user.ID,
		&user.Email,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Image,
		&user.EmailVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, nil, errs.HandleSessionError(err)
	}

	return session, user, nil
}
