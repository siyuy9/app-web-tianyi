package cli

import (
	"github.com/spf13/cobra"
	infraApp "gitlab.com/kongrentian-group/tianyi/v1/infrastructure/app"
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
		infraApp.New().Lifecycle.Run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.AddCommand(serverRun)
}
