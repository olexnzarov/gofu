package commands

import (
	"github.com/olexnzarov/gofu/internal/gofu_cli"
	"github.com/olexnzarov/gofu/pb"
	"github.com/spf13/cobra"
)

var stopCommand = &cobra.Command{
	Use:   "stop {NAME|PID}",
	Short: "Stop a process",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		client, err := gofu_cli.Client()
		if err != nil {
			return err
		}

		reply, err := client.ProcessManager.Stop(
			gofu_cli.Timeout(RequestTimeout),
			&pb.StopRequest{
				Process: args[0],
			},
		)

		if err != nil {
			return gofu_cli.InternalError("failed to stop the process", err)
		}

		if reply.GetError() != nil {
			return gofu_cli.Error("failed to stop the process", reply.GetError())
		}

		return nil
	},
}
