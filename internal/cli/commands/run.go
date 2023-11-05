package commands

import (
	"fmt"
	"maps"
	"time"

	"github.com/olexnzarov/gofu/internal/cli"
	"github.com/olexnzarov/gofu/internal/cli/constants"
	"github.com/olexnzarov/gofu/internal/cli/formatting"
	"github.com/olexnzarov/gofu/pb"
	"github.com/olexnzarov/gofu/pkg/envfmt"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/durationpb"
)

type runFlags struct {
	name              string
	persist           bool
	environment       []string
	environmentFiles  []string
	restart           bool
	restartMaxRetries uint32
	restartDelay      time.Duration
	workingDirectory  string
}

func Run() *cobra.Command {
	return cli.WithFlags(
		cobra.Command{
			Use:   "run COMMAND [ARGUMENT ...]",
			Short: "Start a process",
			Args:  cobra.MinimumNArgs(1),
		},
		&runFlags{},
		run,
	)
}

func run(ctx *cli.Context, flags *runFlags) {
	environmentMap := envfmt.ToKeyValueMap(flags.environment)
	for _, file := range flags.environmentFiles {
		fileEnvMap, err := envfmt.ReadFile(file)
		if err != nil {
			ctx.Output.Add(
				"error",
				formatting.Error(fmt.Sprintf("failed to read the environment file - %s", file), err),
			)
			return
		}
		maps.Copy(environmentMap, fileEnvMap)
	}

	timeout, cancel := ctx.Timeout()
	defer cancel()

	reply, err := ctx.Client.ProcessManager.Start(
		timeout,
		&pb.StartRequest{
			Configuration: &pb.ProcessConfiguration{
				// It's safe to index the slice like that because cobra validated the arguments beforehand.
				Command:     ctx.Args[0],
				Arguments:   ctx.Args[1:],
				Environment: environmentMap,
				Name:        flags.name,
				Persist:     flags.persist,
				RestartPolicy: &pb.ProcessConfiguration_RestartPolicy{
					AutoRestart: flags.restart,
					MaxRetries:  flags.restartMaxRetries,
					Delay:       durationpb.New(flags.restartDelay),
				},
				WorkingDirectory: flags.workingDirectory,
			},
		},
	)
	if out := formatting.ToError(reply, err); out != nil {
		ctx.Output.Add("error", out)
		return
	}

	ctx.Output.Add(
		"process",
		formatting.Process(reply.GetProcess()),
	)
}

func (f *runFlags) Assign(cmd *cobra.Command) {
	cmd.Flags().StringVarP(
		&f.name,
		constants.FLAG_NAME,
		constants.FLAG_SHORT_NAME,
		"",
		"set the process name (defaults to a randomly generated one)",
	)

	cmd.Flags().BoolVarP(
		&f.persist,
		constants.FLAG_PERSIST,
		constants.FLAG_SHORT_PERSIST,
		false,
		"start the process on system startup",
	)

	cmd.Flags().StringArrayVarP(
		&f.environment,
		constants.FLAG_ENV_VALUE,
		constants.FLAG_SHORT_ENV_VALUE,
		[]string{},
		"set an environment variable, usage: -e FOO=BAR -e HELLO=WORLD",
	)

	cmd.Flags().StringArrayVar(
		&f.environmentFiles,
		constants.FLAG_ENV_FILE,
		[]string{},
		"read environment variables from a file, usage: --env-file default.env --env-file local.env",
	)

	cmd.Flags().BoolVarP(
		&f.restart,
		constants.FLAG_RESTART,
		constants.FLAG_SHORT_RESTART,
		false,
		"automatically restart a process when it exits",
	)

	cmd.Flags().Uint32Var(
		&f.restartMaxRetries,
		constants.FLAG_RESTART_MAX_RETRIES,
		0,
		"max number of restart tries",
	)

	cmd.Flags().DurationVar(
		&f.restartDelay,
		constants.FLAG_RESTART_DELAY,
		time.Second*1,
		"delay between automatic restarts",
	)

	cmd.Flags().StringVar(
		&f.workingDirectory,
		constants.FLAG_WORKING_DIRECTORY,
		"",
		"set working directory for the process",
	)
}
