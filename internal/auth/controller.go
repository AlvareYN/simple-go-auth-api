package auth

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	r.Group("/auth")
	{
		r.POST("/login", Login)
		r.POST("/refresh", RefreshToken)
	}
}
