// database setup
package database

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"gitlab.com/kongrentian-group/tianyi/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// connect to the database
func Connect(config *entity.ConfigDatabase) *gorm.DB {
	var dialector gorm.Dialector
	connection := parseConnection(config)
	switch config.Type {
	case entity.ConfigDatabaseTypePostgresql:
		dialector = postgres.Open(connection)
	default:
		log.Panicf("database type '%s' is not supported", config.Type)
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
	return database
}

func parseConnection(config *entity.ConfigDatabase) string {
	buffer := new(bytes.Buffer)
	format := getFormat(config.Type)
	for key, value := range config.Connection {
		fmt.Fprintf(buffer, format, key, value)
	}
	return buffer.String()
}

func getFormat(databaseType entity.ConfigDatabaseType) string {
	switch databaseType {
	case entity.ConfigDatabaseTypePostgresql:
		return "%s=%s "
	}
	return ""
}
