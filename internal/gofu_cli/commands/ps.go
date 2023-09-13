package commands

import (
	"github.com/olexnzarov/gofu/internal/gofu_cli"
	"github.com/olexnzarov/gofu/internal/gofu_cli/outputs"
	"github.com/olexnzarov/gofu/pb"
	"github.com/olexnzarov/gofu/pkg/output"
	"github.com/spf13/cobra"
)

var psCommand = &cobra.Command{
	Use:   "ps",
	Short: "List processes",
	Run: gofu_cli.Run(func(output *output.Output, cmd *cobra.Command, args []string) {
		client, err := gofu_cli.Client()
		if err != nil {
			output.Add(
				"error",
				outputs.Fatal(err),
			)
			return
		}

		reply, err := client.ProcessManager.List(
			gofu_cli.Timeout(RequestTimeout),
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
