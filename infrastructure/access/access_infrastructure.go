package infraAccess

import (
	"log"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	usecaseAccess "gitlab.com/kongrentian-group/tianyi/v1/usecase/access"
	"gorm.io/gorm"
)

// https://github.com/casbin/casbin
// https://github.com/casdoor/casdoor
// https://casbin.org/docs/en/adapters
// https://github.com/casbin/gorm-adapter
type service struct {
	enforcer   *casbin.Enforcer
	database   *gorm.DB
	configPath string
	adapter    *gormadapter.Adapter
}

// https://github.com/casbin/gorm-adapter#customize-table-columns-example
type casbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:512;uniqueIndex:unique_index"`
	V0    string `gorm:"size:512;uniqueIndex:unique_index"`
	V1    string `gorm:"size:512;uniqueIndex:unique_index"`
	V2    string `gorm:"size:512;uniqueIndex:unique_index"`
	V3    string `gorm:"size:512;uniqueIndex:unique_index"`
	V4    string `gorm:"size:512;uniqueIndex:unique_index"`
	V5    string `gorm:"size:512;uniqueIndex:unique_index"`
}

func NewService(
	database *gorm.DB,
	configPath string,
) usecaseAccess.Interactor {
	return &service{database: database, configPath: configPath}
}

func (s *service) SetupEnforcer() {
	gormadapter.TurnOffAutoMigrate(s.database)
	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(
		s.database, &casbinRule{},
	)
	s.adapter = adapter
	if err != nil {
		log.Panicln("could not create the adapter: ", err)
	}
	enforcer, err := casbin.NewEnforcer(s.configPath, adapter)
	if err != nil {
		log.Panicln("could not create the enforcer: ", err)
	}
	s.enforcer = enforcer
	if err = s.Load(); err != nil {
		log.Panicln("could not load the policy: ", err)
	}
}

func (s *service) Migrate() error {
	return s.database.AutoMigrate(&casbinRule{})
}

func (s *service) Load() error {
	return s.enforcer.LoadPolicy()
}

func (s *service) Enforce(subject, object, action interface{}) (
	bool, error,
) {
	return s.enforcer.Enforce(subject, object, action)
}
