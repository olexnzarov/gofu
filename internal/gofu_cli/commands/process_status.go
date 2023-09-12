package commands

import (
	"github.com/olexnzarov/gofu/internal/gofu_cli"
	"github.com/olexnzarov/gofu/internal/gofu_cli/format"
	"github.com/olexnzarov/gofu/pb"
	"github.com/spf13/cobra"
)

var processStatusCommand = &cobra.Command{
	Use:     "ps",
	Aliases: []string{"list"},
	Short:   "List processes",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		client, err := gofu_cli.Client()
		if err != nil {
			return err
		}

		reply, err := client.ProcessManager.List(
			gofu_cli.Timeout(RequestTimeout),
			&pb.ListRequest{},
		)

		if err != nil {
			return gofu_cli.InternalError("failed to list the processes", err)
		}

		format.PrintProcesses(reply.Processes)

		return nil
	},
}
