package model

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ConfigType struct {
	Server   *ServerType   `json:"server"`
	Database *DatabaseType `json:"database"`
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
	err = viper.Unmarshal(config)
	if err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
		return
	}
	log.Fatalf(
		"Could not unmarshal config file %s: %s",
		viper.ConfigFileUsed(),
		err)
	return
}

type ServerType struct {
	Port string `json:"port"`
	Host string `json:"host"`
	// Max number of recent connections during `Expiration` seconds (1 minute)
	// before sending a 429 response
	RequestLimit int `json:"request_limit"`
}

func NewServer() *ServerType {
	return &ServerType{
		Host:         "localhost",
		Port:         "8080",
		RequestLimit: 100,
	}
}

type DatabaseType struct {
	// https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
	Connection map[string]string `json:"connection"`
	Type       DatabaseConstType `json:"type"`
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
