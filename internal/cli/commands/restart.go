package commands

import (
	"github.com/olexnzarov/gofu/internal/cli"
	"github.com/olexnzarov/gofu/internal/cli/formatting"
	"github.com/olexnzarov/gofu/pb"
	"github.com/spf13/cobra"
)

func Restart() *cobra.Command {
	return cli.WithContext(
		cobra.Command{
			Use:   "restart {NAME|PID}",
			Short: "Restart a process",
			Args:  cobra.MinimumNArgs(1),
		},
		restart,
	)
}

func restart(ctx *cli.Context) {
	timeout, cancel := ctx.Timeout()
	defer cancel()

	reply, err := ctx.Client.ProcessManager.Restart(
		timeout,
		&pb.RestartRequest{
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
