package outputs

import (
	"fmt"
	"strings"

	"github.com/olexnzarov/gofu/pb"
)

type ProcessDescriptionOutput struct {
	process *pb.ProcessInformation
}

func ProcessDescription(process *pb.ProcessInformation) *ProcessDescriptionOutput {
	return &ProcessDescriptionOutput{process: process}
}

func (pdo *ProcessDescriptionOutput) getArguments() string {
	if len(pdo.process.Configuration.Arguments) == 0 {
		return "none"
	}
	arguments := strings.Builder{}
	for _, arg := range pdo.process.Configuration.Arguments {
		arguments.WriteString(fmt.Sprintf("\n  %s", arg))
	}
	return arguments.String()
}

func (pdo *ProcessDescriptionOutput) getEnvironment() string {
	if len(pdo.process.Configuration.Environment) == 0 {
		return "none"
	}
	environment := strings.Builder{}
	for key, value := range pdo.process.Configuration.Environment {
		environment.WriteString(fmt.Sprintf("\n  %s: %s", key, value))
	}
	return environment.String()
}

func (pdo *ProcessDescriptionOutput) getRestartPolicy() string {
	policy := pdo.process.Configuration.RestartPolicy
	maxRetries := ""
	if policy.MaxRetries == 0 {
		maxRetries = "none"
	} else {
		maxRetries = fmt.Sprintf("%d", policy.MaxRetries)
	}
	return fmt.Sprintf(
		"\n"+
			"  Enabled: %t\n"+
			"  Delay: %s\n"+
			"  Max retries: %s",
		policy.AutoRestart,
		policy.Delay.AsDuration(),
		maxRetries,
	)
}

func (pdo *ProcessDescriptionOutput) getStatus() string {
	status := strings.ToUpper(pdo.process.Status)
	if pdo.process.ExitState != nil {
		status = fmt.Sprintf(
			"%s\n"+
				"  Exit Code: %d\n"+
				"  Exited At: %s",
			status,
			pdo.process.ExitState.Code,
			pdo.process.ExitState.ExitedAt.AsTime(),
		)
	}
	return status
}

func (pdo *ProcessDescriptionOutput) Text() string {
	process := pdo.process
	config := process.Configuration

	environment := strings.Builder{}
	if len(config.Environment) > 0 {
		for key, value := range config.Environment {
			environment.WriteString(fmt.Sprintf("\n  %s: %s", key, value))
		}
	} else {
		environment.WriteString("none")
	}

	return fmt.Sprintf(
		"Process (%s)\n"+
			"\n--- Configuration ---\n\n"+
			"Name: %s\n"+
			"Persistable: %t\n"+
			"Working directory: %s\n"+
			"Restart policy: %s\n"+
			"Command: %s\n"+
			"Arguments: %s\n"+
			"Environment: %s\n"+
			"\n--- Runtime ---\n\n"+
			"Pid: %d\n"+
			"Started At: %s\n"+
			"Status: %s\n"+
			"Restarts: %d\n"+
			"Logs: %s",
		process.Id,
		config.Name,
		config.Persist,
		config.WorkingDirectory,
		pdo.getRestartPolicy(),
		config.Command,
		pdo.getArguments(),
		pdo.getEnvironment(),
		process.Pid,
		process.StartedAt.AsTime(),
		pdo.getStatus(),
		process.Restarts,
		process.Stdout,
	)
}

func (pdo *ProcessDescriptionOutput) Object() interface{} {
	return pdo.process
}
