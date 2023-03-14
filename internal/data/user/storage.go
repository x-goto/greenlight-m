package user

import (
	"context"
	"goto/greenlight-m/internal/data/user/userdtos"
)

type Repository interface {
	Create(ctx context.Context, user *userdtos.UserCreateDTO) error
	Update(ctx context.Context, user *userdtos.UserUpdateDTO) error
	DeleteByID(ctx context.Context, UserID int) error
	GetByID(ctx context.Context, UserID int) (*userdtos.UserGetDTO, error)
}
