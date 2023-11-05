package cli

import (
	"time"

	"github.com/olexnzarov/gofu/internal/cli/formatting"
	"github.com/spf13/cobra"
)

type Config struct {
	Timeout time.Duration
	Format  string
}

var globalConfig *Config = &Config{}

func InitializeConfig(root *cobra.Command) {
	root.CompletionOptions.HiddenDefaultCmd = true

	root.PersistentFlags().DurationVar(&globalConfig.Timeout, "timeout", time.Second*90, "timeout for requests to the daemon")
	root.PersistentFlags().StringVarP(&globalConfig.Format, "output", "o", formatting.OutputText, "output format (text, json, or prettyjson)")
}
