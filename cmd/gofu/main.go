package main

import (
	"os"

	"github.com/olexnzarov/gofu/internal/cli/commands"
)

func main() {
	root := commands.Root()
	exitCode := commands.Execute(root)
	os.Exit(exitCode)
}
