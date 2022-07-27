// snake_case keys in nested structs does not work if you use `json`,
// you need to use `mapstructure`
// https://github.com/spf13/viper/issues/125
package model

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ConfigType struct {
	Server   *ServerType   `mapstructure:"server"`
	Database *DatabaseType `mapstructure:"database"`
}

type ServerType struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
	// Max number of recent connections during `Expiration` seconds (1 minute)
	// before sending a 429 response
	RequestLimit int           `mapstructure:"request_limit"`
	Fiber        *fiber.Config `mapstructure:"fiber"`
}

type DatabaseType struct {
	// https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
	Connection map[string]string `mapstructure:"connection"`
	Type       DatabaseConstType `mapstructure:"type"`
}

func NewServer() *ServerType {
	return &ServerType{
		Host:         "localhost",
		Port:         "8080",
		RequestLimit: 100,
		Fiber:        &fiber.Config{},
	}
}

func NewDatabase() *DatabaseType {
	return &DatabaseType{
		Type: PostgresqlConst,
		Connection: map[string]string{
			"host":     "localhost",
			"port":     "9920",
			"user":     "admin",
			"password": "admin",
			"dbname":   "tianyi",
			"sslmode":  "disable",
		},
	}
}

func NewConfigType() *ConfigType {
	return &ConfigType{
		Server:   NewServer(),
		Database: NewDatabase(),
	}
}

// reads in config file and ENV variables if set.
func (config *ConfigType) ReadFromEnvironment() {
	// enable automatic load of environment variables that match
	// e.g., TIANYI_CONFIG for '--config' flag
	viper.SetEnvPrefix("tianyi")
	viper.AutomaticEnv()

	if fromCmd := viper.GetString("config"); fromCmd != "" {
		// get config path from cmd or from environment
		viper.SetConfigFile(fromCmd)
	} else {
		// default config paths
		home, error := os.UserHomeDir()
		cobra.CheckErr(error)
		paths := []string{
			".",
			home,
			"/etc/tianyi",
		}
		for _, path := range paths {
			viper.AddConfigPath(path)
		}
		viper.SetConfigType("yml")
		viper.SetConfigName("tianyi")
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error initializing config file:", err)
		return
	}
	err = viper.UnmarshalExact(config)
	if err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
		return
	}
	log.Fatalf(
		"Could not unmarshal config file %s: %s",
		viper.ConfigFileUsed(),
		err)
}

// Config is a struct holding the server settings.
type Config struct {
	// Enables the "Server: value" HTTP header.
	//
	// Default: ""
	ServerHeader string `mapstructure:"server_header"`
}
