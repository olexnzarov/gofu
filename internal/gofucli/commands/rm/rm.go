package rm

import (
	"fmt"

	"github.com/olexnzarov/gofu/internal/gofucli/constants"
	"github.com/olexnzarov/gofu/internal/gofucli/utilities"
	"github.com/olexnzarov/gofu/internal/gofucli/utilities/outputs"
	"github.com/olexnzarov/gofu/internal/output"
	"github.com/olexnzarov/gofu/pb"
	"github.com/spf13/cobra"
)

var (
	force     bool
	alwaysYes bool
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

		if !alwaysYes {
			confirmed := utilities.Confirm(
				fmt.Sprintf(
					"This command will remove the process '%s'. You will not be able to recover it after it is done.",
					args[0],
				),
			)
			if !confirmed {
				return
			}
		}

		timeout, cancel := utilities.Timeout()
		defer cancel()

		reply, err := client.ProcessManager.Remove(
			timeout,
			&pb.RemoveRequest{
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

func init() {
	Command.Flags().BoolVarP(
		&force,
		constants.FLAG_FORCE,
		constants.FLAG_SHORT_FORCE,
		false,
		"stop the process if it's running",
	)

	Command.Flags().BoolVarP(
		&alwaysYes,
		constants.FLAG_ALWAYS_YES,
		constants.FLAG_SHORT_ALWAYS_YES,
		false,
		"do not ask confirmations",
	)
}
