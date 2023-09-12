package commands

import (
	"fmt"
	"time"

	"github.com/olexnzarov/gofu/internal/gofu_cli"
	"github.com/olexnzarov/gofu/pb"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/durationpb"
)

var (
	name              string
	save              bool
	environment       []string
	environmentFiles  []string
	restart           bool
	restartMaxRetries uint32
	restartDelay      time.Duration
	cwd               string
)

var runCommand = &cobra.Command{
	Use:   "run COMMAND [ARGUMENT ...]",
	Short: "Start a process",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		client, err := gofu_cli.Client()
		if err != nil {
			return err
		}

		reply, err := client.ProcessManager.Start(
			gofu_cli.Timeout(RequestTimeout),
			&pb.StartRequest{
				Configuration: &pb.ProcessConfiguration{
					// It's safe to index the slice like that because cobra validated the arguments beforehand.
					Command:   args[0],
					Arguments: args[1:],
					Name:      name,
					Persist:   save,
					RestartPolicy: &pb.ProcessConfiguration_RestartPolicy{
						AutoRestart: restart,
						MaxRetries:  restartMaxRetries,
						Delay:       durationpb.New(restartDelay),
					},
					WorkingDirectory: cwd,
				},
			},
		)

		if err != nil {
			return gofu_cli.InternalError("failed to start the process", err)
		}

		if reply.GetError() != nil {
			return gofu_cli.Error("failed to start the process", reply.GetError())
		}

		fmt.Printf("pid=%d\n", reply.GetProcess().Pid)

		return nil
	},
}

func init() {
	runCommand.Flags().StringVarP(&name, "name", "n", "", "assign a process name (defaults to a randomly generated one)")
	runCommand.Flags().BoolVarP(&save, "save", "s", false, "save the process to run it on startup")
	runCommand.Flags().StringArrayVarP(&environment, "env", "e", []string{}, "set an environment variable, usage: -e FOO=BAR -e HELLO=WORLD")
	runCommand.Flags().StringArrayVar(&environmentFiles, "env-file", []string{}, "read environment variables from a file, usage: --env-file default.env --env-file local.env")
	runCommand.Flags().BoolVarP(&restart, "restart", "r", true, "restart the process when it exits")
	runCommand.Flags().Uint32Var(&restartMaxRetries, "restart-max-retries", 1, "max number of restart tries")
	runCommand.Flags().DurationVar(&restartDelay, "restart-delay", 0, "delay between automatic restarts")
	runCommand.Flags().StringVar(&cwd, "cwd", "", "sets current working directory for the process")
}