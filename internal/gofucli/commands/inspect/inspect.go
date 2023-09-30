package inspect

import (
	"github.com/olexnzarov/gofu/internal/gofucli/utilities"
	"github.com/olexnzarov/gofu/internal/gofucli/utilities/outputs"
	"github.com/olexnzarov/gofu/internal/output"
	"github.com/olexnzarov/gofu/pb"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "inspect {NAME|PID}",
	Short: "Inspect a process",
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

		reply, err := client.ProcessManager.Get(
			timeout,
			&pb.GetRequest{
				Process: args[0],
			},
		)
		if out := outputs.ToError(reply, err); out != nil {
			output.Add("error", out)
			return
		}

		output.Add(
			"process",
			outputs.ProcessDescription(reply.GetProcess()),
		)
	}),
}
