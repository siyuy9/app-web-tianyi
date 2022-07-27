package db

import (
	"bytes"
	"fmt"
	"log"

	"gitlab.com/kongrentian-groups/golang/tianyi/backend/model"
)

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
