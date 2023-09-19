package commands

import (
	"github.com/olexnzarov/gofu/internal/gofu_cli/commands/ps"
	"github.com/olexnzarov/gofu/internal/gofu_cli/commands/restart"
	"github.com/olexnzarov/gofu/internal/gofu_cli/commands/rm"
	"github.com/olexnzarov/gofu/internal/gofu_cli/commands/run"
	"github.com/olexnzarov/gofu/internal/gofu_cli/commands/stop"
	"github.com/olexnzarov/gofu/internal/gofu_cli/commands/update"
	"github.com/spf13/cobra"
)

func Include(root *cobra.Command) {
	root.AddCommand(run.Command)
	root.AddCommand(restart.Command)
	root.AddCommand(stop.Command)
	root.AddCommand(ps.Command)
	root.AddCommand(rm.Command)
	root.AddCommand(update.Command)
}
