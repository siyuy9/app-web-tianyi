package db

import (
	"log"

	"gitlab.com/kongrentian-groups/golang/tianyi/backend/model"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/model/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// database is nil untill you call connect
type DatabaseService struct {
	database *gorm.DB
	config   *model.DatabaseType
}

func NewDatabaseService(config *model.DatabaseType) *DatabaseService {
	return &DatabaseService{
		database: nil,
		config:   config,
	}
}

// connect to database and automigrate
func (service *DatabaseService) Connect() *DatabaseService {
	var dialector gorm.Dialector
	connection := parseConnection(service.config)
	switch service.config.Type {
	case model.PostgresqlConst:
		dialector = postgres.Open(connection)
	default:
		notSupported(service.config.Type)
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
	service.database = database
	return service
}

func (service *DatabaseService) Find(
	dest interface{},
	conds ...interface{},
) *DatabaseService {
	service.database.Find(dest, conds...)
	return service
}
