package userpwds

import (
	"context"
	"goto/greenlight-m/internal/data/user_passwords/userpassworddtos"
)

type Repository interface {
	Create(ctx context.Context, password userpassworddtos.UserPasswordDTO) error
	Update(ctx context.Context, password userpassworddtos.UserPasswordDTO) error
	DeleteByUserID(ctx context.Context, userID string) error
	GetByUserID(ctx context.Context, userID string) (userpassworddtos.UserPasswordDTO, error)
}
