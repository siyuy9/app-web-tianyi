// snake_case keys in nested structs does not work if you use `json`,
// you need to use `mapstructure`
// https://github.com/spf13/viper/issues/125
package model

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/redis"
)

type ConfigType struct {
	Server   *ServerType   `mapstructure:"server"`
	Database *DatabaseType `mapstructure:"database"`
	Redis    *redis.Config `mapstructure:"redis"`
	Fiber    *fiber.Config `mapstructure:"fiber"`
}

type ServerType struct {
	// server port
	Port string `mapstructure:"port"`
	// server host
	Host string `mapstructure:"host"`
	// Max number of recent connections during `Expiration` seconds (1 minute)
	// before sending a 429 response
	RequestLimit int `mapstructure:"request_limit"`
}

type DatabaseType struct {
	// https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
	Connection map[string]string `mapstructure:"connection"`
	Type       DatabaseConstType `mapstructure:"type"`
}
