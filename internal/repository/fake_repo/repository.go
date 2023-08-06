package fakerepo

import (
	"context"

	"github.com/kozhamseitova/test-task/internal/entity"
	"github.com/kozhamseitova/test-task/utils"
)

type Repository struct {
	users      map[int]*entity.CreateUserRequest
	currentID  int
}

func NewRepository() *Repository {
	return &Repository{
		users:      make(map[int]*entity.CreateUserRequest),
		currentID:  1,
	}
}

func (f *Repository) CreateUser(ctx context.Context, u *entity.CreateUserRequest) (int, error) {
	u.Id = f.currentID

	f.currentID++

	f.users[u.Id] = u

	return u.Id, nil
}

func (f *Repository) GetUserByUsername(ctx context.Context, username string) (*entity.CreateUserRequest, error) {
	for _, u := range f.users {
		if u.Username == username {
			return u, nil
		}
	}

	return nil, utils.ErrNotFound
}

func (f *Repository) UpdateUser(ctx context.Context, u *entity.CreateUserRequest) error {
	f.users[u.Id] = u

	return nil
}

func (f *Repository) DeleteUser(ctx context.Context, id int) error {
	delete(f.users, id)

	return nil
}

func (f *Repository) GetAllUsers(ctx context.Context, filter entity.UserFilter) ([]*entity.User, error) {
	var users []*entity.User
	for _, u := range f.users {
		users = append(users, &entity.User{
			Id:        u.Id,
			Username:  u.Username,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			City:      u.City,
			BirthDate: u.BirthDate,
		})
	}

	return users, nil
}

func (f *Repository) GetUsersById(ctx context.Context, id int) (*entity.User, error) {
	u, ok := f.users[id]
	if !ok {
		return nil, utils.ErrNotFound
	}

	return &entity.User{
		Id:        u.Id,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		City:      u.City,
		BirthDate: u.BirthDate,
	}, nil
}
