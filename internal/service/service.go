package service

import (
	"context"

	"github.com/kozhamseitova/test-task/internal/entity"
)

type Service interface {
	CreateUser(ctx context.Context, u *entity.CreateUserRequest) (int, error)
	Login(ctx context.Context, username, password string) (string, error)
	UpdateUser(ctx context.Context, u *entity.CreateUserRequest) error
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context, filter entity.UserFilter) ([]*entity.User, error)
	GetUsersById(ctx context.Context, id int) (*entity.User, error)
	VerifyToken(ctx context.Context, token string) (int, error) 
}