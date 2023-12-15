package main

import (
	"net/http"

	"github.com/AlvareYN/auth-api-go/cmd"
	"github.com/AlvareYN/auth-api-go/internal"
	"github.com/AlvareYN/auth-api-go/internal/auth"
	"github.com/AlvareYN/auth-api-go/internal/users"
	"github.com/gin-gonic/gin"
)

func main() {
	cmd.Setup()
	db := cmd.New()
	r := gin.Default()
	internal.SetupModels(db)

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	users.Routes(&r.RouterGroup)
	auth.Routes(&r.RouterGroup)

	r.Run(":8000")
}
