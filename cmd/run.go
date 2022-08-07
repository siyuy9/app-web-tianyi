package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/kongrentian-group/tianyi/api/controller"
	"gitlab.com/kongrentian-group/tianyi/entity"
	"gitlab.com/kongrentian-group/tianyi/infrastructure/database"
	"gitlab.com/kongrentian-group/tianyi/infrastructure/registry"
	"gitlab.com/kongrentian-group/tianyi/infrastructure/router"
	"gitlab.com/kongrentian-group/tianyi/infrastructure/ui/web2"
	"gitlab.com/kongrentian-group/tianyi/usecase/interactor"
)

func RunServer() {
	config := readConfig()
	fiberRouter := router.New(config.Fiber)
	database := database.Connect(config.Database)
	validator := validator.New()
	sessionRegistry := session.New(session.Config{
		Storage: redis.New(*config.Redis),
	})

	app := &controller.AppController{
		Frontend: controller.NewFrontendController(
			web2.FrontendFilesystem,
		),
		User: controller.NewUserController(
			interactor.NewUserInteractor(
				registry.NewUserRegistry(database),
			),
			validator,
		),
		Lifecycle: controller.NewLifecycleContoller(
			fiberRouter,
			sessionRegistry,
		),
	}

	router.Setup(fiberRouter, app)

	go app.Lifecycle.Listen()
	app.Lifecycle.ShutdownOnInterruptionSignal()
}

// populate the config file
// panic if it is not found or on any other error
func readConfig() *entity.Config {
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
	config, configPath := entity.NewConfig(), viper.ConfigFileUsed()
	if err := viper.UnmarshalExact(config); err != nil {
		log.Panicf("Could not unmarshal config file %s: %s", configPath, err)
	}
	log.Println("Using config file: ", configPath)
	return config
}
