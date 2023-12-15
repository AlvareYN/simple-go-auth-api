package middlewares

import (
	"log"

	"github.com/AlvareYN/auth-api-go/internal/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {

	// Get the JWT from the header
	tokenString := c.Request.Header.Get("Authorization")

	// Validate the token
	claims, err := utils.ValidateToken(tokenString)

	if err != nil {
		utils.HandleException(c, "Invalid token", 401)
		c.Abort()
		return

	}
	log.Println("claims middleware")
	log.Println(claims)
	c.Set("user", claims)

	c.Next()
}
