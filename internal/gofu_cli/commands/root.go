package commands

import (
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	RequestTimeout time.Duration
)

var rootCommand = &cobra.Command{
	Use:   "gofu",
	Short: "gofu - an easy-to-use process manager for all kinds of applications.",
}

func Execute() {
	err := rootCommand.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCommand.PersistentFlags().DurationVar(&RequestTimeout, "timeout", time.Second*90, "timeout for requests to the daemon")

	rootCommand.AddCommand(runCommand)
	rootCommand.AddCommand(restartCommand)
	rootCommand.AddCommand(stopCommand)
	rootCommand.AddCommand(processStatusCommand)
}
