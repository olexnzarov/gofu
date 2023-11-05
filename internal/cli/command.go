package cli

import (
	"fmt"

	"github.com/olexnzarov/gofu/internal/cli/formatting"
	"github.com/spf13/cobra"
)

type FlagAssignable interface {
	comparable
	Assign(*cobra.Command)
}

func onError(err error) {
	HasError = true
	out := formatting.NewOutput()
	out.Add(
		"error",
		formatting.Fatal(err),
	)
	if err := out.Print(globalConfig.Format); err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err))
	}
}

func onExit(ctx *Context) {
	HasError = ctx.Output.HasError
	if err := ctx.Output.Print(globalConfig.Format); err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err))
	}
}

func WithContext(cmd cobra.Command, run func(*Context)) *cobra.Command {
	cmd.Run = func(cmd *cobra.Command, args []string) {
		cmd.SilenceUsage = true
		ctx, err := NewContext(cmd, args)
		if err != nil {
			onError(err)
			return
		}
		run(ctx)
		onExit(ctx)
	}
	return &cmd
}

func WithFlags[T FlagAssignable](cmd cobra.Command, flags T, run func(*Context, T)) *cobra.Command {
	flags.Assign(&cmd)
	cmd.Run = func(cmd *cobra.Command, args []string) {
		ctx, err := NewContext(cmd, args)
		if err != nil {
			onError(err)
			return
		}
		run(ctx, flags)
		onExit(ctx)
	}
	return &cmd
}
