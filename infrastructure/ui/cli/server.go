package cli

import (
	"github.com/spf13/cobra"
	"gitlab.com/kongrentian-group/tianyi/cmd"
)

// serverCmd represents the serve command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run in server mode",
	Long:  `description`,
	Run: func(command *cobra.Command, args []string) {
		cmd.RunServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
