package repository

import (
	"context"
	"database/sql"
	"github.com/linqcod/student-testing-app/app/internal/domain/entity"
)

const (
	CreateTaskCategoryQuery     = "INSERT INTO users (first_name, second_name, email, password) VALUES ($1, $2, $3, $4)"
	DeleteTaskCategoryQuery     = "SELECT * FROM users WHERE id = $1 LIMIT 1"
	GetTaskCategoryQuery        = "SELECT * FROM users"
	GetTaskCategoryByTitleQuery = "UPDATE users SET first_name = $2, second_name = $3, email = $4, password = $5 WHERE id = $1"
	UpdateTaskCategoryQuery     = "DELETE FROM users WHERE id = $1"
	GetAllTaskCategoriesQuery   = "SELECT * FROM users WHERE email = $1 LIMIT 1"
)

type TaskCategoryRepository interface {
	CreateTaskCategory(category *entity.TaskCategory) (*entity.TaskCategory, error)
	DeleteTaskCategory(categoryId int64) error
	GetTaskCategory(categoryId int64) (*entity.TaskCategory, error)
	GetTaskCategoryByTitle(title string) (*entity.TaskCategory, error)
	UpdateTaskCategory(category *entity.TaskCategory) (*entity.TaskCategory, error)
	GetAllTaskCategories() ([]entity.TaskCategory, error)
}

type TaskCategoryRepositoryImpl struct {
	ctx context.Context
	db  *sql.DB
}

func NewTaskCategoryRepository(ctx context.Context, db *sql.DB) TaskCategoryRepository {
	return &TaskCategoryRepositoryImpl{
		ctx: ctx,
		db:  db,
	}
}

func (t TaskCategoryRepositoryImpl) CreateTaskCategory(category *entity.TaskCategory) (*entity.TaskCategory, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskCategoryRepositoryImpl) DeleteTaskCategory(categoryId int64) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskCategoryRepositoryImpl) GetTaskCategory(categoryId int64) (*entity.TaskCategory, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskCategoryRepositoryImpl) GetTaskCategoryByTitle(title string) (*entity.TaskCategory, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskCategoryRepositoryImpl) UpdateTaskCategory(category *entity.TaskCategory) (*entity.TaskCategory, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskCategoryRepositoryImpl) GetAllTaskCategories() ([]entity.TaskCategory, error) {
	//TODO implement me
	panic("implement me")
}
