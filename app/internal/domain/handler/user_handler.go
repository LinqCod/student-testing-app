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

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var registerUser entity.UserRegistrationDTO
	err := c.ShouldBindJSON(&registerUser)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.userService.SaveUser(&registerUser)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseCreated(c, result)
}

func (h *UserHandler) GetUsersList(c *gin.Context) {
	result, err := h.userService.GetUsersList()
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if result == nil {
		result = make([]entity.UserDTO, 0)
	}

	response.ResponseOKWithData(c, result)
}

func (h *UserHandler) GetUserDetails(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.userService.GetUserDetails(int64(userId))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if result == nil {
		result = &entity.UserDTO{}
	}

	response.ResponseOKWithData(c, result)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	var updateUser entity.User
	err = c.ShouldBindJSON(&updateUser)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	updateUser.Id = int64(userId)

	result, err := h.userService.UpdateUser(&updateUser)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if result == nil {
		result = &entity.UserDTO{}
	}

	response.ResponseOKWithData(c, result)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.userService.DeleteUser(int64(userId))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseOK(c, "user deleted successfully")
}

func (h *UserHandler) Login(c *gin.Context) {
	var loginDTO entity.UserLoginDTO

	err := c.ShouldBindJSON(&loginDTO)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	validateUser, err := h.userService.GetUserByEmail(loginDTO)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if validateUser == nil {
		validateUser = &entity.User{}
	}

	token, err := jwttoken.CreateToken(validateUser.Id)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	userData := map[string]interface{}{
		"access_token":    token.AccessToken,
		"expiration_time": token.ExpirationTimeInUnix,
		"user_id":         validateUser.Id,
	}

	response.ResponseOKWithData(c, userData)
}
