package users

import (
	"context"
	"goto/greenlight-m/internal/data/dtos"
)

type Repository interface {
	Create(ctx context.Context, user *dtos.UserRegistrationDTO) error
	GetByID(ctx context.Context, UserID int) (*dtos.UserFetchingDTO, error)
}
