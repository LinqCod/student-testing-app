package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/linqcod/student-testing-app/app/internal/domain/entity"
	"github.com/linqcod/student-testing-app/app/internal/domain/service"
	"github.com/linqcod/student-testing-app/app/pkg/jwttoken"
	"github.com/linqcod/student-testing-app/app/pkg/response"
	"net/http"
	"strconv"
)

type StudentHandler struct {
	studentService service.StudentService
}

func NewStudentHandler(studentService service.StudentService) *StudentHandler {
	return &StudentHandler{
		studentService: studentService,
	}
}

func (h *StudentHandler) RegisterStudent(c *gin.Context) {
	var registerStudent entity.StudentRegistrationDTO
	err := c.ShouldBindJSON(&registerStudent)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.studentService.SaveStudent(&registerStudent)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseCreated(c, result)
}

func (h *StudentHandler) GetStudentDetails(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("student_id"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.studentService.GetStudentDetails(int64(studentId))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseOKWithData(c, result)
}

func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("student_id"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	var updateStudent entity.Student
	err = c.ShouldBindJSON(&updateStudent)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	updateStudent.Id = int64(studentId)

	result, err := h.studentService.UpdateStudent(&updateStudent)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseOKWithData(c, result)
}

func (h *StudentHandler) Login(c *gin.Context) {
	var loginDTO entity.StudentLoginDTO

	err := c.ShouldBindJSON(&loginDTO)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	validateStudent, err := h.studentService.GetStudentByEmail(loginDTO)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := jwttoken.CreateToken(validateStudent.Id)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	studentData := map[string]interface{}{
		"access_token":    token.AccessToken,
		"expiration_time": token.ExpirationTimeInUnix,
		"student_id":      validateStudent.Id,
	}

	response.ResponseOKWithData(c, studentData)
}
