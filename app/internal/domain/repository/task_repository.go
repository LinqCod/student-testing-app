package repository

import (
	"context"
	"database/sql"
	"github.com/linqcod/student-testing-app/app/internal/domain/entity"
)

const (
	InsertTaskQuery   = "INSERT INTO tasks (category_id, description) VALUES ($1, $2)"
	InsertAnswerQuery = "INSERT INTO task_answers (task_id, text, is_right) VALUES ($1, $2, $3)"
	// GetTasksByCategoryQuery DeleteTaskQuery         = "DELETE task_answers, tasks FROM task_answers LEFT JOIN tasks ON task_answers.task_id = tasks.id WHERE task.id = $1"
	GetTasksByCategoryQuery = ""
	UpdateTaskQuery         = ""
	GetAllTasksQuery        = ""
)

type TaskRepository interface {
	CreateTask(task *entity.Task) (*entity.Task, error)
	DeleteTask(taskId int64) error
	GetTasksByCategoryQuery(category string) ([]entity.Task, error)
	UpdateTask(task *entity.Task) (*entity.Task, error)
	GetAllTasks() ([]entity.Task, error)
}

type TaskRepositoryImpl struct {
	ctx context.Context
	db  *sql.DB
}

func NewTaskRepository(ctx context.Context, db *sql.DB) TaskRepository {
	return &TaskRepositoryImpl{
		ctx: ctx,
		db:  db,
	}
}

func (t TaskRepositoryImpl) CreateTask(task *entity.Task) (*entity.Task, error) {
	err := t.db.QueryRowContext(t.ctx, InsertTaskQuery, task.CategoryId, task.Description).Err()
	if err != nil {
		return nil, err
	}

	for i := range task.Answers {
		err = t.db.QueryRowContext(t.ctx, InsertAnswerQuery, task.Answers[i].TaskId, task.Answers[i].Text, task.Answers[i].IsRight).Err()
		if err != nil {
			return nil, err
		}
	}

	return task, nil
}

func (t TaskRepositoryImpl) DeleteTask(taskId int64) error {
	//TODO: implement deletion of task with answers
	return nil
}

func (t TaskRepositoryImpl) GetTasksByCategoryQuery(category string) ([]entity.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskRepositoryImpl) UpdateTask(task *entity.Task) (*entity.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskRepositoryImpl) GetAllTasks() ([]entity.Task, error) {
	//TODO implement me
	panic("implement me")
}
