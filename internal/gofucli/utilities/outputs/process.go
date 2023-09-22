package outputs

import (
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/olexnzarov/gofu/internal/gofudaemon/procmanager"
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

func prettyProcessStatus(status string) string {
	color := text.FgRed
	switch status {
	case procmanager.STATUS_RUNNING:
		color = text.FgGreen
	case procmanager.STATUS_RESTARTING, procmanager.STATUS_STOPPED:
		color = text.FgYellow
	}
	return color.Sprint(strings.ToUpper(status))
}

func processComandWithArguments(process *pb.ProcessInformation) string {
	return strings.Join(append([]string{process.Configuration.Command}, process.Configuration.Arguments...), " ")
}
