package main

import (
	"os"

	"github.com/olexnzarov/gofu/internal/gofu_cli"
)

func main() {
	exitCode := gofu_cli.Execute()
	os.Exit(exitCode)
}
