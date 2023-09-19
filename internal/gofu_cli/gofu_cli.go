package gofu_cli

import (
	"github.com/olexnzarov/gofu/internal/gofu_cli/commands"
	"github.com/olexnzarov/gofu/internal/gofu_cli/utilities"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "gofu",
	Short: "gofu - an easy-to-use process manager for all kinds of applications",
}

func Execute() int {
	err := rootCommand.Execute()
	if err != nil {
		return 1
	}
	return utilities.GetExitCode()
}

func init() {
	utilities.InitializeRoot(rootCommand)
	commands.Include(rootCommand)
}
