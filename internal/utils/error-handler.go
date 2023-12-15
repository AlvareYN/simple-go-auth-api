package utils

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, msg string, status int) {
	log.Println(msg)
	c.JSON(status, gin.H{"status": status, "message": msg})
}

func HandleException(c *gin.Context, msg string, status int) {
	c.JSON(status, gin.H{"status": status, "message": msg})
	return
}

func IsUniqueConstraintError(err error) bool {
	if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		return true
	}
	return false
}
