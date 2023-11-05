package commands

import (
	"github.com/olexnzarov/gofu/internal/cli"
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	root := &cobra.Command{
		Use:   "gofu",
		Short: "gofu - an easy-to-use process manager for all kinds of applications",
	}
	cli.InitializeConfig(root)
	root.AddCommand(
		Inspect(),
		List(),
		Remove(),
		Restart(),
		Run(),
		Stop(),
		Update(),
	)
	return root
}

func Execute(root *cobra.Command) int {
	err := root.Execute()
	if err != nil || cli.HasError {
		return 1
	}
	return 0
}
