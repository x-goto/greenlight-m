package userauthdbs

import (
	"context"
	"database/sql"
	"errors"
	"goto/greenlight-m/internal/data/user_auth/userauthdtos"
	"goto/greenlight-m/pkg/utils/sqlutils"
)

type PQUserAuthDataRepository struct {
	DB *sql.DB
}

func (r *PQUserAuthDataRepository) Create(ctx context.Context, authData userauthdtos.UserAuthDataDTO) error {
	query := `INSERT INTO user_auth_data (user_id, is_activated, activation_link)
				 VALUES ($1, $2, $3)
				 RETURING user_id;`
	err := r.DB.QueryRowContext(ctx, query, authData.UserID, authData.IsActivated, authData.ActivationLink).Scan(authData.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (r *PQUserAuthDataRepository) Update(ctx context.Context, authData userauthdtos.UserAuthDataDTO) error {
	query := `UPDATE user_auth_data
				 SET user_id = $1, is_activated = $2, activation_link = $3
				 WHERE user_id = $1
				 RETURNING user_id;`
	return r.DB.QueryRowContext(ctx, query, authData.UserID, authData.IsActivated, authData.ActivationLink).Scan(authData.UserID)
}

func (r *PQUserAuthDataRepository) DeleteByUserID(ctx context.Context, userID string) error {
	query := `DELETE FROM user_auth_data WHERE user_id = $1;`
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

func (r *PQUserAuthDataRepository) GetByUserID(ctx context.Context, userID string) (*userauthdtos.UserAuthDataDTO, error) {
	query := `SELECT user_id, is_activated, activation_link FROM user_auth_data WHERE user_id = $1;`

	var authData userauthdtos.UserAuthDataDTO

	err := r.DB.QueryRowContext(ctx, query, userID).Scan(
		&authData.UserID,
		&authData.IsActivated,
		&authData.ActivationLink,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, sqlutils.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &authData, nil
}
