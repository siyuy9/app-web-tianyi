package cli

import (
	"encoding/json"
	"log"

	"github.com/spf13/cobra"
	"gitlab.com/kongrentian-group/tianyi/v1/entity"
	infraApp "gitlab.com/kongrentian-group/tianyi/v1/infrastructure/app"
)

var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "database commands",
	Long:  `description`,
}

var databaseMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "migrate the database",
	Long:  "description",
	Run: func(command *cobra.Command, args []string) {
		infraApp.New().Lifecycle.Migrate()
	},
}

var databaseCreate = &cobra.Command{
	Use:   "create",
	Short: "create resources",
	Long:  "description",
}

var databaseAddUser = &cobra.Command{
	Use:   "user",
	Short: "create user",
	Long:  "description",
	Run: func(command *cobra.Command, args []string) {
		interactor := infraApp.New()
		user := &entity.User{
			Username: username,
			Password: password,
			Admin:    admin,
			Email:    email,
		}
		if err := interactor.User.Create(user); err != nil {
			panic(err)
		}
		bytes, err := json.Marshal(user)
		if err != nil {
			panic(err)
		}
		log.Println(string(bytes))
	},
}

var (
	username string
	password string
	admin    = false
	email    string
)

func init() {
	rootCmd.AddCommand(databaseCmd)
	databaseCmd.AddCommand(databaseMigrate)
	databaseCmd.AddCommand(databaseCreate)
	databaseCreate.AddCommand(databaseAddUser)

	databaseAddUser.Flags().StringVarP(
		&username,
		"username",
		"u",
		"",
		"username",
	)
	databaseAddUser.Flags().StringVarP(
		&password,
		"password",
		"p",
		"",
		"password",
	)
	databaseAddUser.Flags().BoolVarP(
		&admin,
		"admin",
		"a",
		false,
		"admin",
	)
	databaseAddUser.Flags().StringVarP(
		&email,
		"email",
		"e",
		"",
		"email",
	)
	err := databaseAddUser.MarkFlagRequired("username")
	if err != nil {
		panic(err)
	}
	err = databaseAddUser.MarkFlagRequired("password")
	if err != nil {
		panic(err)
	}
	err = databaseAddUser.MarkFlagRequired("email")
	if err != nil {
		panic(err)
	}
}
