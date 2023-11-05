package outputs

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/olexnzarov/gofu/internal/daemon/processmanager/service/managedprocess"
	"github.com/olexnzarov/gofu/pb"
)

type ProcessOutput struct {
	process *pb.ProcessInformation
}

func Process(process *pb.ProcessInformation) *ProcessOutput {
	return &ProcessOutput{process: process}
}

func (po *ProcessOutput) Text() string {
	return Processes([]*pb.ProcessInformation{po.process}).Text()
}

func (po *ProcessOutput) Object() interface{} {
	return po.process
}

func prettyProcessStatus(process *pb.ProcessInformation) string {
	color := text.FgRed
	switch process.Status {
	case managedprocess.STATUS_RUNNING:
		color = text.FgGreen
	case managedprocess.STATUS_RESTARTING, managedprocess.STATUS_STOPPED:
		color = text.FgYellow
	}
	status := color.Sprint(strings.ToUpper(process.Status))
	if process.Status == managedprocess.STATUS_RESTARTING {
		if process.Configuration.RestartPolicy.MaxRetries == 0 {
			status = fmt.Sprintf("%s (always)", status)
		} else {
			status = fmt.Sprintf("%s (%d/%d)", status, process.Restarts, process.Configuration.RestartPolicy.MaxRetries)
		}
	}
	return status
}

func processComandWithArguments(process *pb.ProcessInformation) string {
	return strings.Join(append([]string{process.Configuration.Command}, process.Configuration.Arguments...), " ")
}
