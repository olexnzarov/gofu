package update

import (
	"time"

	"github.com/olexnzarov/gofu/internal/gofucli/constants"
	"github.com/olexnzarov/gofu/internal/gofucli/utilities"
	"github.com/olexnzarov/gofu/internal/gofucli/utilities/outputs"
	"github.com/olexnzarov/gofu/pb"
	"github.com/olexnzarov/gofu/pkg/output"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

var (
	name              string
	persist           bool
	restart           bool
	restartMaxRetries uint32
	restartDelay      time.Duration
	workingDirectory  string
)

var Command = &cobra.Command{
	Use:   "update {NAME|PID}",
	Short: "Update a process",
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

		update := newUpdateBuilder(cmd)
		update.includeFlag(
			constants.FLAG_NAME,
			func(config *pb.ProcessConfiguration) string {
				config.Name = name
				return "name"
			},
		)
		update.includeFlag(
			constants.FLAG_PERSIST,
			func(config *pb.ProcessConfiguration) string {
				config.Persist = persist
				return "persist"
			},
		)
		update.includeFlag(
			constants.FLAG_RESTART,
			func(config *pb.ProcessConfiguration) string {
				config.RestartPolicy.AutoRestart = restart
				return "restart_policy.auto_restart"
			},
		)
		update.includeFlag(
			constants.FLAG_RESTART_MAX_RETRIES,
			func(config *pb.ProcessConfiguration) string {
				config.RestartPolicy.MaxRetries = restartMaxRetries
				return "restart_policy.max_retries"
			},
		)
		update.includeFlag(
			constants.FLAG_RESTART_DELAY,
			func(config *pb.ProcessConfiguration) string {
				config.RestartPolicy.Delay = durationpb.New(restartDelay)
				return "restart_policy.delay"
			},
		)
		update.includeFlag(
			constants.FLAG_WORKING_DIRECTORY,
			func(config *pb.ProcessConfiguration) string {
				config.WorkingDirectory = workingDirectory
				return "working_directory"
			},
		)

		if len(update.mask) == 0 {
			output.Add(
				"error",
				outputs.Error("", "you must specify at least one flag for the update to be performed"),
			)
			return
		}

		timeout, cancel := utilities.Timeout()
		defer cancel()

		reply, err := client.ProcessManager.Update(
			timeout,
			&pb.UpdateRequest{
				Process:       args[0],
				Configuration: update.config,
				UpdateMask:    &fieldmaskpb.FieldMask{Paths: update.mask},
			},
		)
		if err != nil {
			output.Add(
				"error",
				outputs.Error("failed to update the process", err),
			)
			return
		}
		if reply.GetError() != nil {
			output.Add(
				"error",
				outputs.Error("failed to update the process", reply.GetError()),
			)
			return
		}

		output.Add(
			"message",
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
		"update the name of a process",
	)

	Command.Flags().BoolVarP(
		&persist,
		constants.FLAG_PERSIST,
		constants.FLAG_SHORT_PERSIST,
		false,
		"update whether a process should be started on system startup",
	)

	Command.Flags().BoolVarP(
		&restart,
		constants.FLAG_RESTART,
		constants.FLAG_SHORT_RESTART,
		true,
		"update whether a process should be automatically restarted",
	)

	Command.Flags().Uint32Var(
		&restartMaxRetries,
		constants.FLAG_RESTART_MAX_RETRIES,
		1,
		"update the max number of restart tries",
	)

	Command.Flags().DurationVar(
		&restartDelay,
		constants.FLAG_RESTART_DELAY,
		0,
		"update the delay between automatic restarts",
	)

	Command.Flags().StringVar(
		&workingDirectory,
		constants.FLAG_WORKING_DIRECTORY,
		"",
		"update the working directory of a process",
	)
}
