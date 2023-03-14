package dbs

import (
	"context"
	"database/sql"
	"errors"
	"goto/greenlight-m/internal/data/user_profiles/userprofiledtos"
	"goto/greenlight-m/pkg/utils/sqlutils"
)

type PQProfilesRepository struct {
	DB *sql.DB
}

func (r *PQProfilesRepository) Create(ctx context.Context, profile userprofiledtos.UserProfileDTO) error {
	query := `INSERT INTO user_profiles (user_id, user_img_url, firstname, lastname, bio)
					 VALUES ($1, $2, $3, $4, $5)
					 RETURNING user_id;`
	err := r.DB.QueryRowContext(ctx, query, profile.UserID, profile.UserImageUrl, profile.FirstName, profile.LastName, profile.Bio).Scan(&profile.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (r *PQProfilesRepository) Update(ctx context.Context, profile userprofiledtos.UserProfileDTO) error {
	query := `UPDATE user_profiles
				 SET user_img_url = $2, firstname = $3, lastname = $4, bio = $5
				 WHERE user_id = $1
				 RETURNING user_id;`
	return r.DB.QueryRowContext(ctx, query, profile.UserID, profile.UserImageUrl, profile.FirstName, profile.LastName, profile.Bio).Scan(&profile.UserID)
}

func (r *PQProfilesRepository) DeleteByUserID(ctx context.Context, userID string) error {
	query := `DELETE FROM user_profiles WHERE user_id = $1`

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

func (r *PQProfilesRepository) GetByUserID(ctx context.Context, userID string) (*userprofiledtos.UserProfileDTO, error) {
	query := `SELECT user_id, user_img_url, firstname, lastname, bio FROM user_profiles
				 WHERE user_id = $1`
	var profile userprofiledtos.UserProfileDTO

	err := r.DB.QueryRowContext(ctx, query, userID).Scan(
		&profile.UserID,
		&profile.UserImageUrl,
		&profile.FirstName,
		&profile.LastName,
		&profile.Bio,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, sqlutils.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &profile, nil
}
