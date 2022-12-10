package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/linqcod/student-testing-app/app/internal/domain/service"
	"github.com/linqcod/student-testing-app/app/pkg/response"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (t TaskHandler) GetTaskCategoriesBySubjectId(c *gin.Context) {
	subjectId, err := strconv.Atoi(c.Param("subject_id"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	categories, err := t.taskService.GetTaskCategoriesBySubjectId(int64(subjectId))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseOKWithData(c, categories)
}

func (t TaskHandler) GetTasksByCategoryId(c *gin.Context) {
	categoryId, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	tasks, err := t.taskService.GetTasksByCategoryId(int64(categoryId))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseOKWithData(c, tasks)
}

func (t TaskHandler) GetTaskAnswersByTaskId(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	taskAnswers, err := t.taskService.GetTaskAnswersByTaskId(int64(taskId))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseOKWithData(c, taskAnswers)
}
