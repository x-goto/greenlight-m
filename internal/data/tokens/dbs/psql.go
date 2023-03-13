package dbs

import (
	"context"
	"database/sql"
	"errors"
	"goto/greenlight-m/internal/data/tokens/tokendtos"
	"goto/greenlight-m/pkg/utils/sqlutils"
)

type PQTokenRepository struct {
	DB *sql.DB
}

func (r *PQTokenRepository) Create(ctx context.Context, token tokendtos.TokenDTO) error {
	query := `INSERT INTO tokens (user_id, refresh_token)
				 VALUES ($1, $2)
				 RETURNING user_id;`
	err := r.DB.QueryRowContext(ctx, query, token.UserID, token.RefreshToken).Scan(&token.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (r *PQTokenRepository) Update(ctx context.Context, token tokendtos.TokenDTO) error {
	query := `UPDATE tokens
				 SET user_id = $1, refresh_token = $2
				 WHERE user_id = $1
				 RETURNING user_id;`
	return r.DB.QueryRowContext(ctx, query, token.UserID, token.RefreshToken).Scan(&token.UserID)
}

func (r *PQTokenRepository) DeleteByUserID(ctx context.Context, UserID string) error {
	query := `DELETE FROM tokens WHERE user_id = $1;`
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

func (r *PQTokenRepository) GetByUserID(ctx context.Context, UserID string) (*tokendtos.TokenDTO, error) {
	query := `SELECT user_id, refresh_token FROM tokens WHERE user_id = $1;`
	var token tokendtos.TokenDTO

	err := r.DB.QueryRowContext(ctx, query, UserID).Scan(
		&token.UserID,
		&token.RefreshToken,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, sqlutils.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &token, nil
}
