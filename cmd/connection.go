package cmd

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const DSN string = "host=localhost user=postgres password=example dbname=authApi port=5432 sslmode=disable"

var db *gorm.DB

func setupGormLogger() {

}

func New() *gorm.DB {
	if db == nil {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second,  // Slow SQL threshold
				LogLevel:                  logger.Error, // Log level
				IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      true,         // Don't include params in the SQL log
				Colorful:                  false,        // Disable color
			},
		)
		conn, err := gorm.Open(postgres.New(postgres.Config{
			DSN: DSN,
		}), &gorm.Config{SkipDefaultTransaction: true, Logger: newLogger})

		if err != nil {
			panic(err)
		}
		return conn
	}
	return db
}
