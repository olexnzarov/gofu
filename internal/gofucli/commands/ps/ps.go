package ps

import (
	"github.com/olexnzarov/gofu/internal/gofucli/utilities"
	"github.com/olexnzarov/gofu/internal/gofucli/utilities/outputs"
	"github.com/olexnzarov/gofu/pb"
	"github.com/olexnzarov/gofu/pkg/output"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:     "ps",
	Short:   "List processes",
	Aliases: []string{"list"},
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

		reply, err := client.ProcessManager.List(
			timeout,
			&pb.ListRequest{},
		)
		if err != nil {
			output.Add(
				"error",
				outputs.Error("failed to get a list of processes", err),
			)
			return
		}

		output.Add(
			"processes",
			outputs.Processes(reply.Processes),
		)
	}),
}
