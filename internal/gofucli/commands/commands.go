package commands

import (
	"github.com/olexnzarov/gofu/internal/gofucli/commands/ps"
	"github.com/olexnzarov/gofu/internal/gofucli/commands/restart"
	"github.com/olexnzarov/gofu/internal/gofucli/commands/rm"
	"github.com/olexnzarov/gofu/internal/gofucli/commands/run"
	"github.com/olexnzarov/gofu/internal/gofucli/commands/stop"
	"github.com/olexnzarov/gofu/internal/gofucli/commands/update"
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
