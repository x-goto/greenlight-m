package user

import (
	"context"
	userdto "goto/greenlight-m/internal/data/user/userdtos"
)

type Repository interface {
	Create(ctx context.Context, user *userdto.CreateUserDTO) error
	UpdateUser(ctx context.Context, user *userdto.UpdateUserDTO) error
	DeleteByID(ctx context.Context, UserID int) error
	GetByID(ctx context.Context, UserID int) (*userdto.GetUserDTO, error)
}
