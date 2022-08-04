package common

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/redis"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/model"
)

func newServerConfig() *model.ServerType {
	return &model.ServerType{
		Host:         "127.0.0.1",
		Port:         "8080",
		RequestLimit: 100,
	}
}

func newDatabaseConfig() *model.DatabaseType {
	return &model.DatabaseType{
		Type: model.PostgresqlConst,
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

func newRedisConfig() *redis.Config {
	return &redis.Config{
		Host:     "127.0.0.1",
		Port:     6379,
		Password: "admin",
		Database: 0,
		Reset:    false,
	}
}

func newAppConfig() *model.ConfigType {
	return &model.ConfigType{
		Server:   newServerConfig(),
		Database: newDatabaseConfig(),
		Redis:    newRedisConfig(),
		Fiber:    &fiber.Config{},
	}
}

// reads in config file and ENV variables if set.
func fillConfigFromEnvironment(config *model.ConfigType) {
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
