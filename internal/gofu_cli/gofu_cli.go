package gofu_cli

import (
	"context"
	"fmt"
	"time"

	"github.com/olexnzarov/gofu/pkg/gofu"
	"github.com/olexnzarov/gofu/pkg/output"
	"github.com/spf13/cobra"
)

type runFunc = func(output *output.Output, cmd *cobra.Command, args []string)
type cobraRunFunc = func(cmd *cobra.Command, args []string)

func Run(run runFunc) cobraRunFunc {
	return func(cmd *cobra.Command, args []string) {
		cmd.SilenceUsage = true

		out := output.NewOutput()
		outputFormat := cmd.Flag("output").Value.String()

		run(out, cmd, args)

		if err := out.Print(outputFormat); err != nil {
			fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		}
	}
}

func Client() (*gofu.Client, error) {
	client, err := gofu.DefaultClient()
	if err != nil {
		return nil, fmt.Errorf("this error may indicate that the gofu daemon is not running, %s", err.Error())
	}
	return client, nil
}

func Timeout(timeout time.Duration) (context.Context, context.CancelFunc) {
	context, cancel := context.WithTimeout(context.Background(), timeout)
	return context, cancel
}
