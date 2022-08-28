package infraConfig

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/redis"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// server config
// snake_case keys in nested structs do not work if you use `json`,
// you need to use `mapstructure`
// https://github.com/spf13/viper/issues/125
type App struct {
	Server   *Server       `mapstructure:"server"`
	Database *Database     `mapstructure:"database"`
	Redis    *redis.Config `mapstructure:"redis"`
	Fiber    *fiber.Config `mapstructure:"fiber"`
}

type Server struct {
	// server port
	Port string `mapstructure:"port"`
	// server host
	Host string `mapstructure:"host"`
	// Max number of recent connections during `Expiration` seconds (1 minute)
	// before sending a 429 response
	RequestLimit int  `mapstructure:"request_limit"`
	JWT          *JWT `mapstructure:"jwt"`
}

type JWT struct {
	// secret signs jwt claims
	Secret string `mapstructure:"secret"`
	// expiration affects how fast jwt claims expire (hours)
	Expiration int `mapstructure:"expiration" validate:"required,min=1"`
}

// jwt middleware requires an array of bytes
func (configJWT *JWT) GetSecret() []byte {
	return []byte(configJWT.Secret)
}

type Database struct {
	// https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
	Connection map[string]string `mapstructure:"connection"`
	Type       DatabaseType      `mapstructure:"type"`
}

type DatabaseType string

const DatabaseTypePostgresql = "postgresql"

func NewServer() *Server {
	return &Server{
		Host:         "127.0.0.1",
		Port:         "8080",
		RequestLimit: 100,
		JWT:          NewJWT(),
	}
}

func NewJWT() *JWT {
	return &JWT{
		Secret:     "",
		Expiration: 72,
	}
}

func NewDatabase() *Database {
	return &Database{
		Type: DatabaseTypePostgresql,
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

func NewRedis() *redis.Config {
	return &redis.Config{
		Host:     "127.0.0.1",
		Port:     6379,
		Password: "admin",
		Database: 0,
		Reset:    false,
	}
}

func New() *App {
	return &App{
		Server:   NewServer(),
		Database: NewDatabase(),
		Redis:    NewRedis(),
		Fiber:    &fiber.Config{},
	}
}

// populate the config file using command line, environment, and a set of
// predefined paths
func (app *App) Populate() *App {
	// enable automatic load of environment variables that match
	// e.g., TIANYI_CONFIG for '--config' flag
	viper.SetEnvPrefix("tianyi")
	viper.AutomaticEnv()

	if fromCmd := viper.GetString("config"); fromCmd != "" {
		// get config path from cmd or from environment
		viper.SetConfigFile(fromCmd)
	} else {
		// default config paths
		homeDirectory, errHome := os.UserHomeDir()
		configDirectory, errConfig := os.UserConfigDir()
		cacheDirectory, errCache := os.UserCacheDir()
		cobra.CheckErr(errHome)
		cobra.CheckErr(errConfig)
		cobra.CheckErr(errCache)
		paths := []string{
			".",
			"tianyi",
			filepath.Join(homeDirectory, "tianyi"),
			filepath.Join(configDirectory, "tianyi"),
			filepath.Join(cacheDirectory, "tianyi"),
			"/opt/tianyi",
		}
		for _, path := range paths {
			viper.AddConfigPath(path)
		}
		viper.SetConfigType("yml")
		viper.SetConfigName("tianyi")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Panicln("Error initializing config file: ", err)
	}
	configPath := viper.ConfigFileUsed()
	if err := viper.UnmarshalExact(app); err != nil {
		log.Panicf(
			"Could not unmarshal the config file %s: %s",
			configPath,
			err,
		)
	}
	log.Println("Using config file: ", configPath)
	return app
}
