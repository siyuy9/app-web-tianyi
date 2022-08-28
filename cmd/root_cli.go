package cli

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "tianyi",
	Short: "tianyi",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP(
		"config",
		"c",
		"",
		"config path",
	)
	err := viper.BindPFlag(
		"config", rootCmd.PersistentFlags().Lookup("config"),
	)
	if err != nil {
		panic(err)
	}
}
