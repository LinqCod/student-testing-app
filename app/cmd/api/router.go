package api

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/linqcod/student-testing-app/app/cmd/api/middleware"
	"github.com/linqcod/student-testing-app/app/internal/domain/handler"
	"github.com/linqcod/student-testing-app/app/internal/domain/repository"
	"github.com/linqcod/student-testing-app/app/internal/domain/service"
)

func InitRouter(ctx context.Context, db *sql.DB) *gin.Engine {
	router := gin.Default()

	userRepository := repository.NewUserRepository(ctx, db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	user := router.Group("v1/user")
	{
		user.GET("/", userHandler.GetUsersList)
		user.GET("/:user_id", middleware.AuthMiddleware(), userHandler.GetUserDetails)
		user.POST("/", userHandler.RegisterUser)
		user.PUT("/:user_id", userHandler.UpdateUser)
		user.DELETE("/:user_id", userHandler.DeleteUser)
		user.POST("/login", userHandler.Login)
	}

	return router
}
