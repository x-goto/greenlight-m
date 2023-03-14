package userauth

import (
	"context"
	"goto/greenlight-m/internal/data/user_auth/userauthdtos"
)

type Repository interface {
	Create(ctx context.Context, authData userauthdtos.UserAuthDataDTO) error
	Update(ctx context.Context, authData userauthdtos.UserAuthDataDTO) error
	DeleteByUserID(ctx context.Context, userID string) error
	GetByUserID(ctx context.Context, userID string) (userauthdtos.UserAuthDataDTO, error)
}
