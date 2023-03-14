package userprofiles

import (
	"context"
	"goto/greenlight-m/internal/data/user_profiles/userprofiledtos"
)

type Repository interface {
	Create(ctx context.Context, profile userprofiledtos.UserProfileDTO) error
	Update(ctx context.Context, profile userprofiledtos.UserProfileDTO) error
	DeleteByUserID(ctx context.Context, userID string) error
	GetByUserID(ctx context.Context, userID string) (userprofiledtos.UserProfileDTO, error)
}
