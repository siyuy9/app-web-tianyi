// snake_case keys in nested structs does not work if you use `json`,
// you need to use `mapstructure`
// https://github.com/spf13/viper/issues/125
package entity

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/redis"
)

// server config
type Config struct {
	Server   *ConfigServer   `mapstructure:"server"`
	Database *ConfigDatabase `mapstructure:"database"`
	Redis    *redis.Config   `mapstructure:"redis"`
	Fiber    *fiber.Config   `mapstructure:"fiber"`
}

type ConfigServer struct {
	// server port
	Port string `mapstructure:"port"`
	// server host
	Host string `mapstructure:"host"`
	// Max number of recent connections during `Expiration` seconds (1 minute)
	// before sending a 429 response
	RequestLimit int `mapstructure:"request_limit"`
}

type ConfigDatabase struct {
	// https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
	Connection map[string]string  `mapstructure:"connection"`
	Type       ConfigDatabaseType `mapstructure:"type"`
}

type ConfigDatabaseType string

const ConfigDatabaseTypePostgresql = "postgresql"

func NewConfigServer() *ConfigServer {
	return &ConfigServer{
		Host:         "127.0.0.1",
		Port:         "8080",
		RequestLimit: 100,
	}
}

func NewConfigDatabase() *ConfigDatabase {
	return &ConfigDatabase{
		Type: ConfigDatabaseTypePostgresql,
		Connection: map[string]string{
			"host":     "127.0.0.1",
			"port":     "9920",
			"user":     "admin",
			"password": "admin",
			"dbname":   "tianyi",
			"sslmode":  "disable",
		},
	}
}

func NewConfigRedis() *redis.Config {
	return &redis.Config{
		Host:     "127.0.0.1",
		Port:     6379,
		Password: "admin",
		Database: 0,
		Reset:    false,
	}
}

func NewConfig() *Config {
	return &Config{
		Server:   NewConfigServer(),
		Database: NewConfigDatabase(),
		Redis:    NewConfigRedis(),
		Fiber:    &fiber.Config{},
	}
}
