/*
Copyright © 2022 kongrentian

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/common"
)

var configFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tianyi",
	Short: "cicd platform",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&configFile,
		"config",
		"c",
		"",
		"config path",
	)
	viper.BindPFlag("config", rootCmd.Flags().Lookup("config"))
	cobra.OnInitialize(common.App.Config.ReadFromEnvironment)
}
