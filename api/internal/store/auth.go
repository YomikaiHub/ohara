package store

import (
	"context"
	"database/sql"

	"github.com/YomikaiHub/ohara/api/internal/errs"
	"github.com/YomikaiHub/ohara/api/internal/models"
)

type AuthStore struct {
	db *sql.DB
}

func (store *AuthStore) RegisterWithEmail(
	ctx context.Context,
	user *models.User,
	account *models.Account,
) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	userQuery := `
		INSERT INTO users (email, first_name, last_name, username, image) VALUES ($1, $2, $3, $4, $5)
		RETURNING id, email, first_name, last_name, username, image, email_verified, created_at, updated_at;
	`

	err = tx.QueryRowContext(
		ctx,
		userQuery,
		user.Email,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Image,
	).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Image,
		&user.EmailVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return errs.HandleUserError(err)
	}

	account.UserID = user.ID

	accountQuery := `
		INSERT INTO accounts (user_id, account_id, provider_id, password_hash) VALUES ($1, $2, $3, $4)
		RETURNING id, account_id, user_id, provider_id, created_at, updated_at;
	`
	err = tx.QueryRowContext(
		ctx,
		accountQuery,
		account.UserID,
		account.AccountID,
		account.ProviderID,
		account.PasswordHash,
	).Scan(
		&account.ID,
		&account.AccountID,
		&account.UserID,
		&account.ProviderID,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err != nil {
		return errs.HandleAccountError(err)
	}

	return tx.Commit()
}

func (store *AuthStore) GetUserWithCredentialByEmail(
	ctx context.Context,
	email string,
) (*models.User, *models.Account, error) {
	query := `
		SELECT
			u.id,
			u.email,
			u.username,
			u.first_name,
			u.last_name,
			u.image,
			u.email_verified,
			u.created_at,
			u.updated_at,
			a.id,
			a.user_id,
			a.account_id,
			a.provider_id,
			a.password_hash,
			a.created_at,
			a.updated_at
		FROM accounts a
		JOIN users u
		ON a.user_id = u.id
		WHERE
		    a.account_id = $1
			AND a.provider_id = 'credential';
	`

	user := &models.User{}
	account := &models.Account{}

	err := store.db.QueryRowContext(
		ctx,
		query,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Image,
		&user.EmailVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
		&account.ID,
		&account.UserID,
		&account.AccountID,
		&account.ProviderID,
		&account.PasswordHash,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err != nil {
		return nil, nil, errs.HandleAccountError(err)
	}

	return user, account, nil
}

func (store *AuthStore) GetUserByID(
	ctx context.Context,
	id string,
) (*models.User, error) {
	query := `
		SELECT
			id,
			email,
			username,
			first_name,
			last_name,
			image,
			email_verified,
			created_at,
			updated_at
		FROM users
		WHERE id = $1;
	`

	user := &models.User{}

	err := store.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
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
		return nil, errs.HandleUserError(err)
	}

	return user, nil
}
