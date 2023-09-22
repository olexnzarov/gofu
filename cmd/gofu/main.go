package main

import (
	"os"

	"github.com/olexnzarov/gofu/internal/gofucli"
)

func main() {
	exitCode := gofucli.Execute()
	os.Exit(exitCode)
}
