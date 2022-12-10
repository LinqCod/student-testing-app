package service

import (
	"github.com/linqcod/student-testing-app/app/internal/domain/entity"
	"github.com/linqcod/student-testing-app/app/internal/domain/repository"
)

type TaskService interface {
	GetTaskCategoriesBySubjectId(subjectId int64) ([]entity.TaskCategory, error)
	GetTasksByCategoryId(categoryId int64) ([]entity.Task, error)
	GetTaskAnswersByTaskId(taskId int64) ([]entity.TaskAnswer, error)
}

type TaskServiceImpl struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(taskRepo repository.TaskRepository) TaskService {
	return &TaskServiceImpl{
		taskRepo: taskRepo,
	}
}

func (t TaskServiceImpl) GetTaskCategoriesBySubjectId(subjectId int64) ([]entity.TaskCategory, error) {
	categories, err := t.taskRepo.GetTaskCategoriesBySubjectId(subjectId)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (t TaskServiceImpl) GetTasksByCategoryId(categoryId int64) ([]entity.Task, error) {
	tasks, err := t.taskRepo.GetTasksByCategoryId(categoryId)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t TaskServiceImpl) GetTaskAnswersByTaskId(taskId int64) ([]entity.TaskAnswer, error) {
	taskAnswers, err := t.taskRepo.GetTaskAnswersByTaskId(taskId)
	if err != nil {
		return nil, err
	}

	return taskAnswers, nil
}
