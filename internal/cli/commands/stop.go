package commands

import (
	"github.com/olexnzarov/gofu/internal/cli"
	"github.com/olexnzarov/gofu/internal/cli/formatting"
	"github.com/olexnzarov/gofu/pb"
	"github.com/spf13/cobra"
)

func Stop() *cobra.Command {
	return cli.WithContext(
		cobra.Command{
			Use:   "stop {NAME|PID}",
			Short: "Stop a process",
			Args:  cobra.MinimumNArgs(1),
		},
		stop,
	)
}

func stop(ctx *cli.Context) {
	timeout, cancel := ctx.Timeout()
	defer cancel()

	reply, err := ctx.Client.ProcessManager.Stop(
		timeout,
		&pb.StopRequest{
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
