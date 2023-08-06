package repository

import (
	"context"

	"github.com/kozhamseitova/test-task/internal/entity"
)

type Repository interface {
	CreateUser(ctx context.Context, u *entity.CreateUserRequest) (int, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.CreateUserRequest, error)
	UpdateUser(ctx context.Context, u *entity.CreateUserRequest) error
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context, filter entity.UserFilter) ([]*entity.User, error)
	GetUsersById(ctx context.Context, id int) (*entity.User, error)
}
