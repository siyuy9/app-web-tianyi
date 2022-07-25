package db

import (
	"bytes"
	"fmt"
	"log"

	"gitlab.com/kongrentian-groups/golang/tianyi/backend/model"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/model/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// connect to database and automigrate
func ConnectDatabase(config *model.ConfigType) *gorm.DB {
	var dialector gorm.Dialector
	connection := parseConnection(config)
	switch config.Database.Type {
	case model.PostgresqlConst:
		dialector = postgres.Open(connection)
	default:
		notSupported(config.Database.Type)
	}

	database, err := gorm.Open(
		dialector,
		&gorm.Config{})
	if err != nil {
		log.Panic("Failed to connect to database: ", err)
	}
	err = database.AutoMigrate(&db.GitlabRunner{})
	if err != nil {
		log.Panic("Failed to migrate database: ", err)
	}
	return database
}

func parseConnection(config *model.ConfigType) string {
	buffer := new(bytes.Buffer)
	format := getFormat(config.Database.Type)
	for key, value := range config.Database.Connection {
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
