package repository

import (
	"database/sql"
	"github.com/linqcod/student-testing-app/app/internal/domain/entity"
)

type UserRepository interface {
	Save(user *entity.User) (*entity.User, error)
	GetDetails(userId int64) (*entity.User, error)
	GetList() ([]entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	Delete(userId int64) error
	GetUserName(userId int64) string
	GetUserByEmailPassword(login entity.UserLoginDTO) (*entity.User, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}
