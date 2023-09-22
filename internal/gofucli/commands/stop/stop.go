package stop

import (
	"github.com/olexnzarov/gofu/internal/gofucli/utilities"
	"github.com/olexnzarov/gofu/internal/gofucli/utilities/outputs"
	"github.com/olexnzarov/gofu/pb"
	"github.com/olexnzarov/gofu/pkg/output"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "stop {NAME|PID}",
	Short: "Stop a process",
	Args:  cobra.MinimumNArgs(1),
	Run: utilities.RunCommand(func(output *output.Output, cmd *cobra.Command, args []string) {
		client, err := utilities.Client()
		if err != nil {
			output.Add(
				"error",
				outputs.Fatal(err),
			)
			return
		}

		timeout, cancel := utilities.Timeout()
		defer cancel()

		reply, err := client.ProcessManager.Stop(
			timeout,
			&pb.StopRequest{
				Process: args[0],
			},
		)
		if err != nil {
			output.Add(
				"error",
				outputs.Error("failed to stop the process", err),
			)
			return
		}
		if reply.GetError() != nil {
			output.Add(
				"error",
				outputs.Error("failed to stop the process", reply.GetError()),
			)
			return
		}

		output.Add(
			"message",
			outputs.Text("OK"),
		)
	}),
}
