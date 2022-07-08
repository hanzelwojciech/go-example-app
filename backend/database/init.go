package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	instance    *gorm.DB
	host        string
	port        string
	username    string
	password    string
	database    string
	initialized = false
)

func CreateDatabaseConnection() (*gorm.DB, error) {
	if !initialized {
		host = os.Getenv("POSTGRES_HOST")
		port = os.Getenv("POSTGRES_PORT")
		username = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		database = os.Getenv("POSTGRES_DATABASE")

		initialized = true
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, username, password, database)

	customLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: customLogger,
	})

	if err != nil {
		return nil, err
	}

	instance = db

	log.Println("[database] Successfully established database connection")

	return instance, nil
}

func GetInstance() *gorm.DB {
	return instance
}
