package commands

import (
	"github.com/olexnzarov/gofu/internal/gofu_cli"
	"github.com/olexnzarov/gofu/internal/gofu_cli/outputs"
	"github.com/olexnzarov/gofu/pb"
	"github.com/olexnzarov/gofu/pkg/output"
	"github.com/spf13/cobra"
)

var restartCommand = &cobra.Command{
	Use:   "restart {NAME|PID}",
	Short: "Restart a process",
	Args:  cobra.MinimumNArgs(1),
	Run: gofu_cli.Run(func(output *output.Output, cmd *cobra.Command, args []string) {
		client, err := gofu_cli.Client()
		if err != nil {
			output.Add(
				"error",
				outputs.Fatal(err),
			)
			return
		}

		reply, err := client.ProcessManager.Restart(
			gofu_cli.Timeout(RequestTimeout),
			&pb.RestartRequest{
				Process: args[0],
			},
		)
		if err != nil {
			output.Add(
				"error",
				outputs.Error("failed to restart the process", err),
			)
			return
		}
		if reply.GetError() != nil {
			output.Add(
				"error",
				outputs.Error("failed to restart the process", reply.GetError()),
			)
			return
		}

		output.Add(
			"message",
			outputs.Text("OK"),
		)
	}),
}
