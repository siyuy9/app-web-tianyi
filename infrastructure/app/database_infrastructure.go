package infraApp

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	infraConfig "gitlab.com/kongrentian-group/tianyi/v1/infrastructure/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func connectDatabase(config *infraConfig.Database) *gorm.DB {
	var dialector gorm.Dialector
	connection := parseConnection(config)
	switch config.Type {
	case infraConfig.DatabaseTypePostgresql:
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

func parseConnection(config *infraConfig.Database) string {
	buffer := &bytes.Buffer{}
	format := getFormat(config.Type)
	for key, value := range config.Connection {
		fmt.Fprintf(buffer, format, key, value)
	}
	return buffer.String()
}

func getFormat(databaseType infraConfig.DatabaseType) string {
	switch databaseType {
	case infraConfig.DatabaseTypePostgresql:
		return "%s=%s "
	}
	return ""
}
