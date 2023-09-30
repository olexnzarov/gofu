package restart

import (
	"github.com/olexnzarov/gofu/internal/gofucli/utilities"
	"github.com/olexnzarov/gofu/internal/gofucli/utilities/outputs"
	"github.com/olexnzarov/gofu/internal/output"
	"github.com/olexnzarov/gofu/pb"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "restart {NAME|PID}",
	Short: "Restart a process",
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

		reply, err := client.ProcessManager.Restart(
			timeout,
			&pb.RestartRequest{
				Process: args[0],
			},
		)
		if out := outputs.ToError(reply, err); out != nil {
			output.Add("error", out)
			return
		}

		output.Add(
			"message",
			outputs.Text("OK"),
		)
	}),
}
