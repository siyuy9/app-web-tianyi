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
type interactor struct {
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

func NewInteractor(
	database *gorm.DB,
	configPath string,
) usecaseAccess.Interactor {
	return &interactor{database: database, configPath: configPath}
}

func (interactor *interactor) SetupEnforcer() {
	gormadapter.TurnOffAutoMigrate(interactor.database)
	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(
		interactor.database,
		&casbinRule{},
	)
	interactor.adapter = adapter
	if err != nil {
		log.Panicln("could not create the adapter: ", err)
	}
	enforcer, err := casbin.NewEnforcer(interactor.configPath, adapter)
	if err != nil {
		log.Panicln("could not create the enforcer: ", err)
	}
	interactor.enforcer = enforcer
	if err = interactor.Load(); err != nil {
		log.Panicln("could not load the policy: ", err)
	}
}

func (interactor *interactor) Migrate() error {
	return interactor.database.AutoMigrate(&casbinRule{})
}

func (interactor *interactor) Load() error {
	return interactor.enforcer.LoadPolicy()
}

func (interactor *interactor) Enforce(subject, object, action interface{}) (
	bool, error,
) {
	return interactor.enforcer.Enforce(subject, object, action)
}
