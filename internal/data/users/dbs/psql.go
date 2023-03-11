package dbs

import (
	"context"
	"database/sql"
	"errors"
	"goto/greenlight-m/internal/data/dtos"
	"goto/greenlight-m/pkg/utils/sqlutils"
)

type PQUserRepository struct {
	DB *sql.DB
}

func (r *PQUserRepository) Create(ctx context.Context, user *dtos.UserRegistrationDTO) error {
	query := `INSERT INTO users (username, email) 
				VALUES ($1, $2)
				RETURNING id;`

	err := r.DB.QueryRowContext(ctx, query, user.Username, user.Email).Scan(&user.ID)

	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return sqlutils.ErrDuplicateEmail
		case err.Error() == `pq: duplicate key value violates unique constraint "users_username_key"`:
			return sqlutils.ErrDuplicateUsername
		default:
			return err
		}
	}

	return nil
}

func (r *PQUserRepository) GetByID(ctx context.Context, UserID int) (*dtos.UserFetchingDTO, error) {
	query := `SELECT id, role, username, email FROM users WHERE id = $1`

	var user dtos.UserFetchingDTO
	err := r.DB.QueryRowContext(ctx, query, UserID).Scan(
		&user.ID,
		&user.Role,
		&user.Username,
		&user.Email,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, sqlutils.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, err
}
