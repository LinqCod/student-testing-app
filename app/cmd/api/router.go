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

	subjectRepository := repository.NewSubjectRepository(ctx, db)
	subjectService := service.NewSubjectService(subjectRepository)
	subjectHandler := handler.NewSubjectHandler(subjectService)

	taskRepository := repository.NewTaskRepository(ctx, db)
	taskService := service.NewTaskService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)

	api := router.Group("/api/v1")
	{
		students := api.Group("/students")
		{
			students.POST("/", studentHandler.RegisterStudent)
			students.POST("/login", studentHandler.Login)
			students.GET("/:student_id", middleware.AuthMiddleware(), studentHandler.GetStudentDetails)
			students.PUT("/:student_id", middleware.AuthMiddleware(), studentHandler.UpdateStudent)
		}
		// FOR STUDENTS
		subjects := api.Group("/subjects")
		{
			subjects.GET("/:group_id", middleware.AuthMiddleware(), subjectHandler.GetSubjectsByGroupId)
		}
		// FOR STUDENTS
		tasks := api.Group("/tasks")
		{
			tasks.GET("/:category_id", middleware.AuthMiddleware(), taskHandler.GetTasksByCategoryId)
			tasks.GET("/categories/:subject_id", middleware.AuthMiddleware(), taskHandler.GetTaskCategoriesBySubjectId)
			tasks.GET("/answers/:task_id", middleware.AuthMiddleware(), taskHandler.GetTaskAnswersByTaskId)
		}
	}

	return router
}
