package outputs

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/olexnzarov/gofu/internal/gofu_cli/format"
	"github.com/olexnzarov/gofu/pb"
)

type ProcessTableOutput struct {
	processes []*pb.ProcessInformation
}

func Processes(processes []*pb.ProcessInformation) *ProcessTableOutput {
	return &ProcessTableOutput{processes: processes}
}

func (pto *ProcessTableOutput) Text() string {
	if len(pto.processes) == 0 {
		return "Found no processes."
	}

	rows := make([]table.Row, 0, len(pto.processes))
	for _, p := range pto.processes {
		rows = append(
			rows,
			table.Row{
				p.Pid,
				p.Configuration.Name,
				format.Truncate(processComandWithArguments(p), 25),
				prettyProcessStatus(p.Status),
			},
		)
	}

	return format.NewTable(
		table.Row{"pid", "name", "command", "status"},
		rows,
	).Render()
}

func (pto *ProcessTableOutput) Object() interface{} {
	if pto.processes == nil {
		return []*pb.ProcessInformation{}
	}
	return pto.processes
}
