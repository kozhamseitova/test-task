package service

import (
	"context"
	"errors"

	"github.com/kozhamseitova/test-task/internal/entity"
	"github.com/kozhamseitova/test-task/pkg/utils"
)

func(m *Manager) CreateUser(ctx context.Context, u *entity.CreateUserRequest) (int, error) {
	user, err := m.repository.GetUserByUsername(ctx, u.Username)
	if err != nil {
		if !errors.Is(err, utils.ErrNotFound) {
			m.logger.Errorf(ctx, "[GetUserByUsername] err: %v", err)
			return 0, utils.ErrInternalError
		}
	}

	if user != nil {
		return 0, utils.ErrUserAlreadyExists
	}

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		m.logger.Errorf(ctx, "[HashPassword] err: %v", err)
		return 0, utils.ErrInternalError
	}

	u.Password = hashedPassword

	id, err := m.repository.CreateUser(ctx, u)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func(m *Manager) Login(ctx context.Context, username, password string) (string, error) {
	user, err := m.repository.GetUserByUsername(ctx, username)

	if err != nil {
		return "", err
	}

	err = utils.CheckPassword(password, user.Password)

	if err != nil {
		m.logger.Errorf(ctx, "[ChechPassword] err: %v", err)
		return "", utils.ErrInternalError
	}

	token, err := m.jwttoken.CreateToken(user.Id, m.config.Token.TimeToLive)

	if err != nil {
		m.logger.Errorf(ctx, "[CreateToken] err: %v", err)
		return "", utils.ErrInternalError
	}

	return token, nil
}

func(m *Manager) UpdateUser(ctx context.Context, u *entity.CreateUserRequest) error {
	return m.repository.UpdateUser(ctx, u)
}

func(m *Manager) DeleteUser(ctx context.Context, id int) error {
	return m.repository.DeleteUser(ctx, id)
}

func(m *Manager) GetAllUsers(ctx context.Context, filter entity.UserFilter) ([]*entity.User, error) {
	return m.repository.GetAllUsers(ctx, filter)
}

func(m *Manager) GetUsersById(ctx context.Context, id int) (*entity.User, error) {
	return m.repository.GetUsersById(ctx, id)
}

func(m *Manager) VerifyToken(ctx context.Context, token string) (int, error) {
	payload, err := m.jwttoken.ValidateToken(token)
	if err != nil {
		m.logger.Errorf(ctx, "[ValidateToken] err: %v", err)
		return 0, utils.ErrInternalError
	}

	return payload.UserId, nil
}
