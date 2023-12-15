package users

import (
	"log"

	"github.com/AlvareYN/auth-api-go/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context) {
	var user UserRequest
	var db *gorm.DB = c.MustGet("db").(*gorm.DB)
	log.Println(c.Request.Body)
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.HandleException(c, "Invalid request body", 400)
		return
	}

	hash := utils.HashPassword(user.Password)
	newUser := User{
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
		Hash:     hash,
		Role:     user.Role,
	}
	log.Println(newUser)
	err := db.Create(&newUser).Error

	if err != nil {
		if utils.IsUniqueConstraintError(err) {
			utils.HandleException(c, "Username or email already exists", 400)
			return
		}
		utils.HandleError(c, "Error creating user", 500)
		return
	}

	c.JSON(201, gin.H{"status": 201, "message": "User created successfully!", "resourceId": newUser.ID})
}

func GetUser(c *gin.Context) {
	var db *gorm.DB = c.MustGet("db").(*gorm.DB)
	var user User
	id := c.Param("id")

	err := db.First(&user, id).Error

	if err != nil {
		utils.HandleError(c, "User not found", 404)
		return
	}

	c.JSON(200, gin.H{"status": 200, "data": user})
}

func GetUsers(c *gin.Context) {
	var db *gorm.DB = c.MustGet("db").(*gorm.DB)
	var users []User

	err := db.Select("id, username, name, email, role").Find(&users).Error

	if err != nil {
		utils.HandleError(c, "Error fetching users", 500)
		return
	}

	c.JSON(200, gin.H{"status": 200, "data": users})
}

func UpdateUser(c *gin.Context) {
	var db *gorm.DB = c.MustGet("db").(*gorm.DB)
	var updatedUser User

	id := c.Param("id")

	var userReq UserRequest

	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Invalid request body"})
	}

	user := User{
		Username: userReq.Username,
		Name:     userReq.Name,
		Email:    userReq.Email,
	}

	if userReq.Password != "" {
		hash := utils.HashPassword(userReq.Password)
		user.Hash = hash
	}
	err := db.Model(&updatedUser).Where("id = ?", id).Updates(user).Error

	if err != nil {
		utils.HandleError(c, "Error updating user", 500)
		return
	}

	c.JSON(200, gin.H{"status": 200, "message": "User updated successfully!"})
}

func DeleteUser(c *gin.Context) {
	var db *gorm.DB = c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var user User

	err := db.First(&user, id).Error

	if err != nil {
		utils.HandleError(c, "User not found", 404)
		return
	}

	err = db.Delete(&user).Error

	if err != nil {
		utils.HandleError(c, "Error deleting user", 500)
		return
	}

	c.JSON(200, gin.H{"status": 200, "message": "User deleted successfully!"})
}
