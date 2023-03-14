package token

import (
	"context"
	"goto/greenlight-m/internal/data/token/tokendtos"
)

type Repository interface {
	Create(ctx context.Context, token tokendtos.TokenDTO) error
	Update(ctx context.Context, token tokendtos.TokenDTO) error
	DeleteByUserID(ctx context.Context, UserID string) error
	GetByUserID(ctx context.Context, UserID string) (*tokendtos.TokenDTO, error)
}
