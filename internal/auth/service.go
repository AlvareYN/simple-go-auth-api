package auth

import (
	"net/http"

	"github.com/AlvareYN/auth-api-go/internal/users"
	"github.com/AlvareYN/auth-api-go/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user users.User
	err := db.Where("email = ?", req.Email).First(&user).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !utils.CheckPassword(req.Password, user.Hash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(map[string]interface{}{"id": user.ID, "email": user.Email, "role": user.Role})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	refreshToken, err := utils.GenerateToken(map[string]interface{}{"id": user.ID, "email": user.Email, "role": user.Role, "refresh": true})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "refresh_token": refreshToken})

}

func RefreshToken(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var req RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, err := utils.ValidateToken(req.RefreshToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	var user users.User
	err = db.Where("email = ?", claims["email"]).First(&user).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(map[string]interface{}{"id": user.ID, "email": user.Email, "role": user.Role})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	refreshToken, err := utils.GenerateToken(map[string]interface{}{"id": user.ID, "email": user.Email, "role": user.Role, "refresh": true})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "refresh_token": refreshToken})

}
