package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/linqcod/student-testing-app/app/pkg/jwttoken"
	"github.com/linqcod/student-testing-app/app/pkg/response"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwttoken.CheckTokenIsValid(c.Request)
		if err != nil {
			response.ResponseError(c, err.Error(), http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Next()
	}
}
