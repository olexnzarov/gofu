package utilities

import (
	"fmt"
	"time"

	"github.com/olexnzarov/gofu/pkg/output"
	"github.com/spf13/cobra"
)

type runFunc = func(output *output.Output, cmd *cobra.Command, args []string)
type cobraRunFunc = func(cmd *cobra.Command, args []string)

var (
	OutputFormat   string
	RequestTimeout time.Duration
)

func InitializeRoot(cmd *cobra.Command) {
	cmd.CompletionOptions.HiddenDefaultCmd = true

	cmd.PersistentFlags().DurationVar(&RequestTimeout, "timeout", time.Second*90, "timeout for requests to the daemon")
	cmd.PersistentFlags().StringVarP(&OutputFormat, "output", "o", output.OutputText, "output format (text, json, or prettyjson)")
}

func RunCommand(run runFunc) cobraRunFunc {
	return func(cmd *cobra.Command, args []string) {
		cmd.SilenceUsage = true

		out := output.NewOutput()
		outputFormat := cmd.Flag("output").Value.String()

		run(out, cmd, args)

		if err := out.Print(outputFormat); err != nil {
			fmt.Println(fmt.Sprintf("Error: %s", err))
		}
	}
}
