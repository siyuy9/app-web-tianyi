package cmd

import (
	"github.com/spf13/cobra"
	infra "gitlab.com/kongrentian-group/tianyi/v1/infrastructure"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server commands",
	Long:  `description`,
}

var serverRun = &cobra.Command{
	Use:   "run",
	Short: "start the server",
	Long:  "description",
	Run: func(command *cobra.Command, args []string) {
		infra.NewApp().Run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.AddCommand(serverRun)
}
