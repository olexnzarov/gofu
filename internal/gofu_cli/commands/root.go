package commands

import (
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	RequestTimeout time.Duration
	OutputFormat   string
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
	rootCommand.CompletionOptions.HiddenDefaultCmd = true

	rootCommand.PersistentFlags().DurationVar(&RequestTimeout, "timeout", time.Second*90, "timeout for requests to the daemon")
	rootCommand.PersistentFlags().StringVarP(&OutputFormat, "output", "o", "text", "output format - text or json")

	rootCommand.AddCommand(runCommand)
	rootCommand.AddCommand(restartCommand)
	rootCommand.AddCommand(stopCommand)
	rootCommand.AddCommand(psCommand)
}
