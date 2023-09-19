package rm

import (
	"github.com/olexnzarov/gofu/internal/gofu_cli/utilities"
	"github.com/olexnzarov/gofu/internal/gofu_cli/utilities/outputs"
	"github.com/olexnzarov/gofu/pb"
	"github.com/olexnzarov/gofu/pkg/output"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:     "rm {NAME|PID}",
	Short:   "Remove a process",
	Aliases: []string{"remove"},
	Args:    cobra.MinimumNArgs(1),
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

		reply, err := client.ProcessManager.Remove(
			timeout,
			&pb.RemoveRequest{
				Process: args[0],
			},
		)
		if err != nil {
			output.Add(
				"error",
				outputs.Error("failed to remove the process", err),
			)
			return
		}
		if reply.GetError() != nil {
			output.Add(
				"error",
				outputs.Error("failed to remove the process", reply.GetError()),
			)
			return
		}

		output.Add(
			"message",
			outputs.Text("OK"),
		)
	}),
}
