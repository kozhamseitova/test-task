package repository

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/kozhamseitova/test-task/internal/entity"
	"github.com/kozhamseitova/test-task/pkg/utils"
)

func (m *Manager) CreateUser(ctx context.Context, u *entity.CreateUserRequest) (int, error) {
	var id int

	query := fmt.Sprintf(`INSERT INTO %s (
						username, password, first_name, last_name, city, birth_date)
						VALUES ($1, $2, $3, $4, $5, $6)
						RETURNING id`, usersTable)

	err := pgxscan.Get(ctx, m.pool, &id, query, u.Username, u.Password, u.FirstName, u.LastName, u.City, u.BirthDate)

	if err != nil {
		m.logger.Errorf(ctx, "[CreateUser] err: %v", err)
		return 0, utils.ErrInternalError
	}

	return id, nil
}

func (m *Manager) GetUserByUsername(ctx context.Context, username string) (*entity.CreateUserRequest, error) {
	var user entity.CreateUserRequest

	query := fmt.Sprintf(`SELECT * from %s WHERE username = $1 LIMIT 1`, usersTable)

	err := pgxscan.Get(ctx, m.pool, &user, query, username)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, utils.ErrNotFound
		}
		m.logger.Errorf(ctx, "[GetUserByUsername] err: %v", err)
		return nil, utils.ErrInternalError
	}

	return &user, nil
}

func (m *Manager) UpdateUser(ctx context.Context, u *entity.CreateUserRequest) error {
	query := fmt.Sprintf(`UPDATE %s SET 
										username=$1,
										password=$2,
										first_name=$3,
										last_name=$4,
										city=$5,
										birth_date=$6
									WHERE id=$7`, usersTable)

	_, err := m.pool.Exec(ctx, query, u.Username, u.Password, u.FirstName, u.LastName, u.City, u.BirthDate, u.Id)

	if err != nil {
		m.logger.Errorf(ctx, "[UpdateUser] err: %v", err)
		return utils.ErrInternalError
	}

	return nil
}

func (m *Manager) DeleteUser(ctx context.Context, id int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, usersTable)

	_, err := m.pool.Exec(ctx, query, id)

	if err != nil {
		m.logger.Errorf(ctx, "[DeleteUser] err: %v", err)
		return utils.ErrInternalError
	}

	return nil
}

func (m *Manager) GetAllUsers(ctx context.Context, filter entity.UserFilter) ([]*entity.User, error) {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Select("id, username, first_name, last_name, city, birth_date").From(usersTable)

	if filter.City != nil {
		builder = builder.Where("city LIKE ?", fmt.Sprint("%", *filter.City, "%"))
	}
	if filter.BirthDateAsc != nil {
		if *filter.BirthDateAsc {
			builder = builder.OrderBy("birth_date ASC")
		} else {
			builder = builder.OrderBy("birth_date DESC")
		}
	}
	if filter.Search != nil {
		searchText := "%" + *filter.Search + " %"
		builder = builder.Where(sq.Or{
			sq.Like{"username": searchText},
			sq.Like{"first_name": searchText},
			sq.Like{"last_name": searchText},
		})
	}

	if (filter.Page > 0 && filter.Amount > 0) {
		offset := filter.Amount * (filter.Page - 1)
		builder = builder.Offset(uint64(offset)).Limit(uint64(filter.Amount))
	}
	

	var filteredUsers []*entity.User
	sql, args, err := builder.ToSql()
	if err != nil {
		m.logger.Errorf(ctx, "[GetAllUsers] err: %v", err)
		return nil, utils.ErrInternalError
	}
	m.logger.Infof(ctx, "sql: %s", sql)

	err = pgxscan.Select(ctx, m.pool, &filteredUsers, sql, args...)
	if err != nil {
		m.logger.Errorf(ctx, "[GetAllUsers] err: %v", err)
		return nil, utils.ErrInternalError
	}

	return filteredUsers, nil
}

func (m *Manager) GetUsersById(ctx context.Context, id int) (*entity.User, error) {
	var user entity.User

	query := fmt.Sprintf(`SELECT * from %s WHERE id = $1 LIMIT 1`, usersTable)

	err := pgxscan.Get(ctx, m.pool, &user, query, id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, utils.ErrNotFound
		}
		m.logger.Errorf(ctx, "[GetUserById] err: %v", err)
		return nil, utils.ErrInternalError
	}

	return &user, nil
}
