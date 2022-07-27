package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend"
)

// serverCmd represents the serve command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run in server mode",
	Long:  `description`,
	Run: func(cmd *cobra.Command, args []string) {
		backend.Start()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	cobra.OnInitialize(backend.Start)
}
