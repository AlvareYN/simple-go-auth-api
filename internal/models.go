package internal

import (
	"github.com/AlvareYN/auth-api-go/internal/users"
	"gorm.io/gorm"
)

func SetupModels(db *gorm.DB) {
	db.AutoMigrate(users.User{})
	// recreate if doent exist
	db.Migrator().DropTable(&users.User{})
	db.Migrator().CreateTable(&users.User{})
}
