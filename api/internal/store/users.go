package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/YomikaiHub/ohara/api/internal/errs"
	"github.com/YomikaiHub/ohara/api/internal/models"
)

type UserStore struct {
	db *sql.DB
}

func (store *UserStore) GetUserByID(
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

func (store *UserStore) Update(
	ctx context.Context,
	userID string,
	username *string,
	firstName *string,
	lastName *string,
	image *string,
) (*models.User, error) {
	var (
		updates []string
		args    []any
	)

	if username != nil {
		updates = append(
			updates,
			fmt.Sprintf("username = $%d", len(args)+1),
		)
		args = append(args, *username)
	}

	if firstName != nil {
		updates = append(
			updates,
			fmt.Sprintf("first_name = $%d", len(args)+1),
		)
		args = append(args, *firstName)
	}

	if lastName != nil {
		updates = append(
			updates,
			fmt.Sprintf("last_name = $%d", len(args)+1),
		)
		args = append(args, lastName)
	}

	if image != nil {
		updates = append(
			updates,
			fmt.Sprintf("image = $%d", len(args)+1),
		)
		args = append(args, image)
	}

	if len(updates) == 0 {
		return nil, errs.ErrNoFieldsToUpdate
	}

	args = append(args, userID)

	query := fmt.Sprintf(`
		UPDATE users
		SET %s,
			updated_at = NOW()
		WHERE id = $%d
		RETURNING
			id,
			email,
			first_name,
			last_name,
			username,
			image,
			email_verified,
			created_at,
			updated_at;
	`, strings.Join(updates, ", "), len(args))

	user := &models.User{}

	err := store.db.QueryRowContext(
		ctx,
		query,
		args...,
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrUserNotFound
		}

		return nil, errs.HandleUserUpdateError(err)
	}

	return user, nil
}
