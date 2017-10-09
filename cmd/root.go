package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	masterCmd.AddCommand(masterStartCmd)
	workerCmd.AddCommand(workerStartCmd)

	RootCmd.AddCommand(masterCmd)
	RootCmd.AddCommand(workerCmd)
}

var RootCmd = &cobra.Command{
	Use:   "prober",
	Short: "prober is a very fast static site generator",
	Long: `A distributed probe framework system used for
				user to detect famous protocols and user defined protocols`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
