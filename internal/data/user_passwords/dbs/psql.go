package dbs

import (
	"context"
	"database/sql"
	"errors"
	"goto/greenlight-m/internal/data/user_passwords/userpassworddtos"
	"goto/greenlight-m/pkg/utils/sqlutils"
)

type PQPasswordRepository struct {
	DB *sql.DB
}

func (r *PQPasswordRepository) Create(ctx context.Context, password userpassworddtos.UserPasswordDTO) error {
	query := `INSERT INTO pwds (user_id, password)
				 VALUES($1, $2)
				 RETURNING user_id;`
	err := r.DB.QueryRowContext(ctx, query, password.UserID, password.Password).Scan(password.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (r *PQPasswordRepository) Update(ctx context.Context, password userpassworddtos.UserPasswordDTO) error {
	query := `UPDATE pwds SET password = $2 WHERE user_id = $1
				 RETURNING user_id;`
	return r.DB.QueryRowContext(ctx, query, password.UserID, password.Password).Scan(password.UserID)
}

func (r *PQPasswordRepository) DeleteByUserID(ctx context.Context, userID string) error {
	query := `DELETE FROM pwds WHERE user_id = $1`

	result, err := r.DB.ExecContext(ctx, query, userID)
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

func (r *PQPasswordRepository) GetByUserID(ctx context.Context, userID string) (*userpassworddtos.UserPasswordDTO, error) {
	query := `SELECT user_id, password FROM pwds WHERE user_id = $1`

	var password userpassworddtos.UserPasswordDTO

	err := r.DB.QueryRowContext(ctx, query, userID).Scan(&password.UserID, &password.Password)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, sqlutils.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &password, nil
}
