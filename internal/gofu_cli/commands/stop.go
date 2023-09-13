package commands

import (
	"github.com/olexnzarov/gofu/internal/gofu_cli"
	"github.com/olexnzarov/gofu/internal/gofu_cli/output"
	"github.com/olexnzarov/gofu/pb"
	"github.com/spf13/cobra"
)

var stopCommand = &cobra.Command{
	Use:   "stop {NAME|PID}",
	Short: "Stop a process",
	Args:  cobra.MinimumNArgs(1),
	Run: gofu_cli.Run(func(output *output.Output, cmd *cobra.Command, args []string) {
		client, err := gofu_cli.Client()
		if err != nil {
			output.Fail(err)
			return
		}

		reply, err := client.ProcessManager.Stop(
			gofu_cli.Timeout(RequestTimeout),
			&pb.StopRequest{
				Process: args[0],
			},
		)
		if err != nil {
			output.Error("failed to stop the process", err)
			return
		}
		if reply.GetError() != nil {
			output.DaemonError("failed to stop the process", reply.GetError())
			return
		}

		output.Text("message", "OK")
	}),
}
