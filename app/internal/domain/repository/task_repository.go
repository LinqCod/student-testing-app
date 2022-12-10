package repository

import (
	"context"
	"database/sql"
	"github.com/linqcod/student-testing-app/app/internal/domain/entity"
)

const (
	GetTaskCategoriesBySubjectIdQuery = "SELECT id, title FROM task_categories WHERE subject_id=$1"
	GetTasksByCategoryIdQuery         = "SELECT id, text FROM tasks WHERE category_id=$1"
	GetTaskAnswersByTaskId            = "SELECT id, text, is_right FROM task_answers WHERE task_id=$1"
)

type TaskRepository interface {
	GetTaskCategoriesBySubjectId(subjectId int64) ([]entity.TaskCategory, error)
	GetTasksByCategoryId(categoryId int64) ([]entity.Task, error)
	GetTaskAnswersByTaskId(taskId int64) ([]entity.TaskAnswer, error)
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

func (t TaskRepositoryImpl) GetTaskCategoriesBySubjectId(subjectId int64) ([]entity.TaskCategory, error) {
	rows, err := t.db.QueryContext(t.ctx, GetTaskCategoriesBySubjectIdQuery, subjectId)
	if err != nil {
		return nil, err
	}

	var categories []entity.TaskCategory
	for rows.Next() {
		var category entity.TaskCategory
		if err = rows.Scan(&category.Id, &category.Title); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (t TaskRepositoryImpl) GetTasksByCategoryId(categoryId int64) ([]entity.Task, error) {
	rows, err := t.db.QueryContext(t.ctx, GetTasksByCategoryIdQuery, categoryId)
	if err != nil {
		return nil, err
	}

	var tasks []entity.Task
	for rows.Next() {
		var task entity.Task
		if err = rows.Scan(&task.Id, &task.Text); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (t TaskRepositoryImpl) GetTaskAnswersByTaskId(taskId int64) ([]entity.TaskAnswer, error) {
	rows, err := t.db.QueryContext(t.ctx, GetTaskAnswersByTaskId, taskId)
	if err != nil {
		return nil, err
	}

	var taskAnswers []entity.TaskAnswer
	for rows.Next() {
		var taskAnswer entity.TaskAnswer
		if err = rows.Scan(&taskAnswer.Id, &taskAnswer.Text, &taskAnswer.IsRight); err != nil {
			return nil, err
		}

		taskAnswers = append(taskAnswers, taskAnswer)
	}

	return taskAnswers, nil
}
