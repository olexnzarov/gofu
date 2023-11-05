package commands

import (
	"fmt"

	"github.com/olexnzarov/gofu/internal/cli"
	"github.com/olexnzarov/gofu/internal/cli/constants"
	"github.com/olexnzarov/gofu/internal/cli/formatting"
	"github.com/olexnzarov/gofu/pb"
	"github.com/spf13/cobra"
)

type removeFlags struct {
	force     bool
	alwaysYes bool
}

func Remove() *cobra.Command {
	return cli.WithFlags(
		cobra.Command{
			Use:     "rm {NAME|PID}",
			Short:   "Remove a process",
			Aliases: []string{"remove"},
			Args:    cobra.MinimumNArgs(1),
		},
		&removeFlags{},
		remove,
	)
}

func remove(ctx *cli.Context, flags *removeFlags) {
	if !flags.alwaysYes {
		confirmed := cli.Confirm(
			fmt.Sprintf(
				"This command will remove the process '%s'. You will not be able to recover it after it is done.",
				ctx.Args[0],
			),
		)
		if !confirmed {
			return
		}
	}

	timeout, cancel := ctx.Timeout()
	defer cancel()

	reply, err := ctx.Client.ProcessManager.Remove(
		timeout,
		&pb.RemoveRequest{
			Process: ctx.Args[0],
		},
	)
	if out := formatting.ToError(reply, err); out != nil {
		ctx.Output.Add("error", out)
		return
	}

	ctx.Output.Add(
		"message",
		formatting.Text("OK"),
	)
}

func (f *removeFlags) Assign(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(
		&f.force,
		constants.FLAG_FORCE,
		constants.FLAG_SHORT_FORCE,
		false,
		"stop the process if it's running",
	)
	cmd.Flags().BoolVarP(
		&f.alwaysYes,
		constants.FLAG_ALWAYS_YES,
		constants.FLAG_SHORT_ALWAYS_YES,
		false,
		"do not ask confirmations",
	)
}
