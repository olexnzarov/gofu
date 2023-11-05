package commands

import (
	"github.com/olexnzarov/gofu/internal/cli"
	"github.com/olexnzarov/gofu/internal/cli/formatting"
	"github.com/olexnzarov/gofu/pb"
	"github.com/spf13/cobra"
)

func List() *cobra.Command {
	return cli.WithContext(
		cobra.Command{
			Use:     "ps",
			Short:   "List processes",
			Aliases: []string{"list"},
		},
		list,
	)
}

func list(ctx *cli.Context) {
	timeout, cancel := ctx.Timeout()
	defer cancel()

	reply, err := ctx.Client.ProcessManager.List(
		timeout,
		&pb.ListRequest{},
	)
	if err != nil {
		ctx.Output.Add("error", formatting.Fatal(err))
		return
	}

	ctx.Output.Add(
		"processes",
		formatting.Processes(reply.Processes),
	)
}
