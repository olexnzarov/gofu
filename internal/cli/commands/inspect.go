package commands

import (
	"github.com/olexnzarov/gofu/internal/cli"
	"github.com/olexnzarov/gofu/internal/cli/formatting"
	"github.com/olexnzarov/gofu/pb"
	"github.com/spf13/cobra"
)

func Inspect() *cobra.Command {
	return cli.WithContext(
		cobra.Command{
			Use:   "inspect {NAME|PID}",
			Short: "Inspect a process",
			Args:  cobra.MinimumNArgs(1),
		},
		inspect,
	)
}

func inspect(ctx *cli.Context) {
	timeout, cancel := ctx.Timeout()
	defer cancel()

	reply, err := ctx.Client.ProcessManager.Get(
		timeout,
		&pb.GetRequest{
			Process: ctx.Args[0],
		},
	)
	if out := formatting.ToError(reply, err); out != nil {
		ctx.Output.Add("error", out)
		return
	}

	ctx.Output.Add(
		"process",
		formatting.ProcessDescription(reply.GetProcess()),
	)
}
