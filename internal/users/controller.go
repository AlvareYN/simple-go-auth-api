package users

import (
	"github.com/AlvareYN/auth-api-go/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {

	r.POST("/users", middlewares.AuthMiddleware, CreateUser)
	r.GET("/users", middlewares.AuthMiddleware, GetUsers)
	r.GET("/users/:id", middlewares.AuthMiddleware, GetUser)
	r.PUT("/users/:id", middlewares.AuthMiddleware, UpdateUser)
	r.DELETE("/users/:id", middlewares.AuthMiddleware, DeleteUser)
}
