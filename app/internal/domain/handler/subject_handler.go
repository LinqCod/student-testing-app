package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/linqcod/student-testing-app/app/internal/domain/service"
	"github.com/linqcod/student-testing-app/app/pkg/response"
	"net/http"
	"strconv"
)

type SubjectHandler struct {
	subjectService service.SubjectService
}

func NewSubjectHandler(subjectService service.SubjectService) *SubjectHandler {
	return &SubjectHandler{
		subjectService: subjectService,
	}
}

func (h *SubjectHandler) GetSubjectsByGroupId(c *gin.Context) {
	groupId, err := strconv.Atoi(c.Param("group_id"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	subjects, err := h.subjectService.GetSubjectsByGroupId(int64(groupId))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseOKWithData(c, subjects)
}
