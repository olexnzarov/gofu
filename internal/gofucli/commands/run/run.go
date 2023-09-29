package run

import (
	"fmt"
	"maps"
	"time"

	"github.com/olexnzarov/gofu/internal/gofucli/constants"
	"github.com/olexnzarov/gofu/internal/gofucli/utilities"
	"github.com/olexnzarov/gofu/internal/gofucli/utilities/outputs"
	"github.com/olexnzarov/gofu/internal/output"
	"github.com/olexnzarov/gofu/pb"
	"github.com/olexnzarov/gofu/pkg/envfmt"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/durationpb"
)

var (
	name              string
	persist           bool
	environment       []string
	environmentFiles  []string
	restart           bool
	restartMaxRetries uint32
	restartDelay      time.Duration
	workingDirectory  string
)

var Command = &cobra.Command{
	Use:   "run COMMAND [ARGUMENT ...]",
	Short: "Start a process",
	Args:  cobra.MinimumNArgs(1),
	Run: utilities.RunCommand(func(output *output.Output, cmd *cobra.Command, args []string) {
		client, err := utilities.Client()
		if err != nil {
			output.Add(
				"error",
				outputs.Fatal(err),
			)
			return
		}

		environmentMap := envfmt.ToKeyValueMap(environment)
		for _, file := range environmentFiles {
			fileEnvMap, err := envfmt.ReadFile(file)
			if err != nil {
				output.Add(
					"error",
					outputs.Error(fmt.Sprintf("failed to read the environment file - %s", file), err),
				)
				return
			}
			maps.Copy(environmentMap, fileEnvMap)
		}

		timeout, cancel := utilities.Timeout()
		defer cancel()

		reply, err := client.ProcessManager.Start(
			timeout,
			&pb.StartRequest{
				Configuration: &pb.ProcessConfiguration{
					// It's safe to index the slice like that because cobra validated the arguments beforehand.
					Command:     args[0],
					Arguments:   args[1:],
					Environment: environmentMap,
					Name:        name,
					Persist:     persist,
					RestartPolicy: &pb.ProcessConfiguration_RestartPolicy{
						AutoRestart: restart,
						MaxRetries:  restartMaxRetries,
						Delay:       durationpb.New(restartDelay),
					},
					WorkingDirectory: workingDirectory,
				},
			},
		)
		if err != nil {
			output.Add(
				"error",
				outputs.Error("failed to start the process", err),
			)
			return
		}
		if reply.GetError() != nil {
			output.Add(
				"error",
				outputs.Error("failed to start the process", reply.GetError()),
			)
			return
		}

		output.Add(
			"process",
			outputs.Process(reply.GetProcess()),
		)
	}),
}

func init() {
	Command.Flags().StringVarP(
		&name,
		constants.FLAG_NAME,
		constants.FLAG_SHORT_NAME,
		"",
		"set the process name (defaults to a randomly generated one)",
	)

	Command.Flags().BoolVarP(
		&persist,
		constants.FLAG_PERSIST,
		constants.FLAG_SHORT_PERSIST,
		false,
		"start the process on system startup",
	)

	Command.Flags().StringArrayVarP(
		&environment,
		constants.FLAG_ENV_VALUE,
		constants.FLAG_SHORT_ENV_VALUE,
		[]string{},
		"set an environment variable, usage: -e FOO=BAR -e HELLO=WORLD",
	)

	Command.Flags().StringArrayVar(
		&environmentFiles,
		constants.FLAG_ENV_FILE,
		[]string{},
		"read environment variables from a file, usage: --env-file default.env --env-file local.env",
	)

	Command.Flags().BoolVarP(
		&restart,
		constants.FLAG_RESTART,
		constants.FLAG_SHORT_RESTART,
		false,
		"automatically restart a process when it exits",
	)

	Command.Flags().Uint32Var(
		&restartMaxRetries,
		constants.FLAG_RESTART_MAX_RETRIES,
		0,
		"max number of restart tries",
	)

	Command.Flags().DurationVar(
		&restartDelay,
		constants.FLAG_RESTART_DELAY,
		time.Second*1,
		"delay between automatic restarts",
	)

	Command.Flags().StringVar(
		&workingDirectory,
		constants.FLAG_WORKING_DIRECTORY,
		"",
		"set working directory for the process",
	)
}
