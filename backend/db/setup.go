package db

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"gitlab.com/kongrentian-groups/golang/tianyi/backend/model"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/model/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// connect to database and automigrate
func Connect(config *model.DatabaseType) *gorm.DB {
	var dialector gorm.Dialector
	connection := parseConnection(config)
	switch config.Type {
	case model.PostgresqlConst:
		dialector = postgres.Open(connection)
	default:
		notSupported(config.Type)
	}

	databaseLogger := logger.New(
		log.New(os.Stdout, "\n", log.LstdFlags), // io writer
		logger.Config{
			// Slow SQL threshold
			SlowThreshold: time.Millisecond * 10,
			// Log level
			LogLevel: logger.Info,
			// Ignore ErrRecordNotFound error for logger
			IgnoreRecordNotFoundError: false,
			// enable color
			Colorful: true,
		},
	)
	database, err := gorm.Open(
		dialector,
		&gorm.Config{
			Logger: databaseLogger,
		})
	if err != nil {
		log.Panic("Failed to connect to database: ", err)
	}
	err = database.AutoMigrate(&db.GitlabRunnerModel{})
	if err != nil {
		log.Panic("Failed to migrate database: ", err)
	}
	return database
}

func parseConnection(config *model.DatabaseType) string {
	buffer := new(bytes.Buffer)
	format := getFormat(config.Type)
	for key, value := range config.Connection {
		fmt.Fprintf(buffer, format, key, value)
	}
	return buffer.String()
}

func getFormat(databaseType model.DatabaseConstType) string {
	switch databaseType {
	case model.PostgresqlConst:
		return "%s=%s "
	default:
		notSupported(databaseType)
	}
	return ""
}

func notSupported(databaseType model.DatabaseConstType) {
	log.Panicf(
		"Database type '%s' is not supported",
		databaseType)
}
