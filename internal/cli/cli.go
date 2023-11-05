package cli

import (
	"context"
	"fmt"

	"github.com/olexnzarov/gofu/internal/cli/formatting"
	"github.com/olexnzarov/gofu/pkg/gofu"
	"github.com/spf13/cobra"
)

type Context struct {
	Client  *gofu.Client
	Output  *formatting.Output
	Command *cobra.Command
	Args    []string
}

var HasError = false

func (ctx *Context) Timeout() (context.Context, context.CancelFunc) {
	context, cancel := context.WithTimeout(context.Background(), globalConfig.Timeout)
	return context, cancel
}

func NewContext(command *cobra.Command, args []string) (*Context, error) {
	client, err := gofu.DefaultClient()
	if err != nil {
		return nil, fmt.Errorf("this error may indicate that the gofu daemon is not running, %s", err)
	}
	return &Context{
		Client:  client,
		Command: command,
		Args:    args,
		Output:  formatting.NewOutput(),
	}, nil
}
