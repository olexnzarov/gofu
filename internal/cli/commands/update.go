package commands

import (
	"time"

	"github.com/olexnzarov/gofu/internal/cli"
	"github.com/olexnzarov/gofu/internal/cli/constants"
	"github.com/olexnzarov/gofu/internal/cli/formatting"
	"github.com/olexnzarov/gofu/pb"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

type updateFlags struct {
	name              string
	persist           bool
	restart           bool
	restartMaxRetries uint32
	restartDelay      time.Duration
	workingDirectory  string
}

func Update() *cobra.Command {
	return cli.WithFlags(
		cobra.Command{
			Use:   "update {NAME|PID}",
			Short: "Update a process",
			Args:  cobra.MinimumNArgs(1),
		},
		&updateFlags{},
		update,
	)
}

func update(ctx *cli.Context, flags *updateFlags) {
	update := newUpdateBuilder(ctx.Command)
	update.includeFlag(
		constants.FLAG_NAME,
		func(config *pb.ProcessConfiguration) string {
			config.Name = flags.name
			return "name"
		},
	)
	update.includeFlag(
		constants.FLAG_PERSIST,
		func(config *pb.ProcessConfiguration) string {
			config.Persist = flags.persist
			return "persist"
		},
	)
	update.includeFlag(
		constants.FLAG_RESTART,
		func(config *pb.ProcessConfiguration) string {
			config.RestartPolicy.AutoRestart = flags.restart
			return "restart_policy.auto_restart"
		},
	)
	update.includeFlag(
		constants.FLAG_RESTART_MAX_RETRIES,
		func(config *pb.ProcessConfiguration) string {
			config.RestartPolicy.MaxRetries = flags.restartMaxRetries
			return "restart_policy.max_retries"
		},
	)
	update.includeFlag(
		constants.FLAG_RESTART_DELAY,
		func(config *pb.ProcessConfiguration) string {
			config.RestartPolicy.Delay = durationpb.New(flags.restartDelay)
			return "restart_policy.delay"
		},
	)
	update.includeFlag(
		constants.FLAG_WORKING_DIRECTORY,
		func(config *pb.ProcessConfiguration) string {
			config.WorkingDirectory = flags.workingDirectory
			return "working_directory"
		},
	)

	if len(update.mask) == 0 {
		ctx.Output.Add(
			"error",
			formatting.Fatal("you must specify at least one flag for the update to be performed"),
		)
		return
	}

	timeout, cancel := ctx.Timeout()
	defer cancel()

	reply, err := ctx.Client.ProcessManager.Update(
		timeout,
		&pb.UpdateRequest{
			Process:       ctx.Args[0],
			Configuration: update.config,
			UpdateMask:    &fieldmaskpb.FieldMask{Paths: update.mask},
		},
	)
	if out := formatting.ToError(reply, err); out != nil {
		ctx.Output.Add("error", out)
		return
	}

	ctx.Output.Add(
		"message",
		formatting.Process(reply.GetProcess()),
	)
}

func (f *updateFlags) Assign(cmd *cobra.Command) {
	cmd.Flags().StringVarP(
		&f.name,
		constants.FLAG_NAME,
		constants.FLAG_SHORT_NAME,
		"",
		"update the name of a process",
	)

	cmd.Flags().BoolVarP(
		&f.persist,
		constants.FLAG_PERSIST,
		constants.FLAG_SHORT_PERSIST,
		false,
		"update whether a process should be started on system startup",
	)

	cmd.Flags().BoolVarP(
		&f.restart,
		constants.FLAG_RESTART,
		constants.FLAG_SHORT_RESTART,
		true,
		"update whether a process should be automatically restarted",
	)

	cmd.Flags().Uint32Var(
		&f.restartMaxRetries,
		constants.FLAG_RESTART_MAX_RETRIES,
		1,
		"update the max number of restart tries",
	)

	cmd.Flags().DurationVar(
		&f.restartDelay,
		constants.FLAG_RESTART_DELAY,
		0,
		"update the delay between automatic restarts",
	)

	cmd.Flags().StringVar(
		&f.workingDirectory,
		constants.FLAG_WORKING_DIRECTORY,
		"",
		"update the working directory of a process",
	)
}

type updateBuilder struct {
	cmd    *cobra.Command
	config *pb.ProcessConfiguration
	mask   []string
}

func newUpdateBuilder(cmd *cobra.Command) *updateBuilder {
	return &updateBuilder{
		cmd:  cmd,
		mask: []string{},
		config: &pb.ProcessConfiguration{
			RestartPolicy: &pb.ProcessConfiguration_RestartPolicy{},
		},
	}
}

func (u *updateBuilder) includeFlag(flagName string, updateFunc func(config *pb.ProcessConfiguration) (field string)) {
	flag := u.cmd.Flags().Lookup(flagName)
	if flag == nil || !flag.Changed {
		return
	}
	field := updateFunc(u.config)
	u.mask = append(u.mask, field)
}
