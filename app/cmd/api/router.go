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

	studentRepository := repository.NewStudentRepository(ctx, db)
	studentService := service.NewStudentService(studentRepository)
	studentHandler := handler.NewStudentHandler(studentService)

	students := router.Group("/api/v1/students")
	{
		students.POST("/", studentHandler.RegisterStudent)
		students.POST("/login", studentHandler.Login)
		students.GET("/:student_id", middleware.AuthMiddleware(), studentHandler.GetStudentDetails)
		students.PUT("/:student_id", middleware.AuthMiddleware(), studentHandler.UpdateStudent)
	}

	return router
}
