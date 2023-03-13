package dbs

import (
	"context"
	"database/sql"
	"errors"
	userdto "goto/greenlight-m/internal/data/user/userdtos"
	"goto/greenlight-m/pkg/utils/sqlutils"
)

type PQUserRepository struct {
	DB *sql.DB
}

func (r *PQUserRepository) Create(ctx context.Context, user *userdto.CreateUserDTO) error {
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

func (r *PQUserRepository) Update(ctx context.Context, user *userdto.UpdateUserDTO) error {
	query := `UPDATE users 
				SET email = $1, username = $2
				WHERE id = $3
				RETURNING email;`
	return r.DB.QueryRowContext(ctx, query, user.Email, user.Username, user.ID).Scan(&user.Email)
}

func (r *PQUserRepository) DeleteByID(ctx context.Context, UserID int) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.DB.ExecContext(ctx, query, UserID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sqlutils.ErrRecordNotFound
	}

	return nil
}

func (r *PQUserRepository) GetByID(ctx context.Context, UserID int) (*userdto.GetUserDTO, error) {
	query := `SELECT id, role, username, email FROM users WHERE id = $1`

	var user userdto.GetUserDTO
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

	return &user, nil
}
