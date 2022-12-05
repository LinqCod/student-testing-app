package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/linqcod/student-testing-app/app/internal/domain/entity"
)

const (
	SaveUserQuery       = "INSERT INTO users (first_name, second_name, email, password) VALUES ($1, $2, $3, $4)"
	GetUserDetailsQuery = "SELECT * FROM users WHERE id = $1 LIMIT 1"
	GetUsersListQuery   = "SELECT * FROM users"
	UpdateUserQuery     = "UPDATE users SET first_name = $2, second_name = $3, email = $4, password = $5 WHERE id = $1"
	DeleteUserQuery     = "DELETE FROM users WHERE id = $1"
	GetUserByEmailQuery = "SELECT * FROM users WHERE email = $1 LIMIT 1"
)

type UserRepository interface {
	SaveUser(user *entity.User) (*entity.User, error)
	GetUserDetails(userId int64) (*entity.User, error)
	GetAllUsers() ([]entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
	DeleteUser(userId int64) error
	GetUserFullName(userId int64) (string, error)
	GetUserByEmail(login entity.UserLoginDTO) (*entity.User, error)
}

type UserRepositoryImpl struct {
	ctx context.Context
	db  *sql.DB
}

func NewUserRepository(ctx context.Context, db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		ctx: ctx,
		db:  db,
	}
}

func (r *UserRepositoryImpl) SaveUser(user *entity.User) (*entity.User, error) {
	err := r.db.QueryRowContext(r.ctx, SaveUserQuery, user.FirstName, user.SecondName, user.Email, user.Password).Err()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) GetUserDetails(userId int64) (*entity.User, error) {
	var user entity.User
	if err := r.db.QueryRowContext(r.ctx, GetUserDetailsQuery, userId).Scan(
		&user.Id,
		&user.FirstName,
		&user.SecondName,
		&user.Email,
		&user.Password,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) GetAllUsers() ([]entity.User, error) {
	var users []entity.User
	rows, err := r.db.QueryContext(r.ctx, GetUsersListQuery)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}

		var id int64
		var firstName string
		var secondName string
		var email string
		var password string
		if err := rows.Scan(&id, &firstName, &secondName, &email, &password); err != nil {
			return nil, err
		}
		users = append(users, entity.User{
			Id:         id,
			FirstName:  firstName,
			SecondName: secondName,
			Email:      email,
			Password:   password,
		})
	}

	return users, nil
}

func (r *UserRepositoryImpl) UpdateUser(user *entity.User) (*entity.User, error) {
	err := r.db.QueryRowContext(r.ctx, UpdateUserQuery, user.Id, user.FirstName, user.SecondName, user.Email, user.Password).Err()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) DeleteUser(userId int64) error {
	err := r.db.QueryRowContext(r.ctx, DeleteUserQuery, userId).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) GetUserFullName(userId int64) (string, error) {
	userDetails, err := r.GetUserDetails(userId)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", userDetails.FirstName, userDetails.SecondName), nil
}

func (r *UserRepositoryImpl) GetUserByEmail(login entity.UserLoginDTO) (*entity.User, error) {
	var user entity.User
	if err := r.db.QueryRowContext(r.ctx, GetUserByEmailQuery, login.Email).Scan(
		&user.Id,
		&user.FirstName,
		&user.SecondName,
		&user.Email,
		&user.Password,
	); err != nil {
		return nil, err
	}

	return &user, nil
}
