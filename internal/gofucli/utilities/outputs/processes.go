package outputs

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/olexnzarov/gofu/internal/gofucli/utilities/format"
	"github.com/olexnzarov/gofu/pb"
)

type ProcessTableOutput struct {
	processes []*pb.ProcessInformation
}

func Processes(processes []*pb.ProcessInformation) *ProcessTableOutput {
	return &ProcessTableOutput{processes: processes}
}

func formatBool(value bool) string {
	if value {
		return "Yes"
	}
	return "No"
}

func (pto *ProcessTableOutput) Text() string {
	if len(pto.processes) == 0 {
		return "Found no processes."
	}

	rows := make([]table.Row, 0, len(pto.processes))
	for _, p := range pto.processes {
		when := "unknown"
		if p.StoppedAt != nil {
			when = fmt.Sprintf("Stopped %s", humanize.Time(p.StoppedAt.AsTime()))
		} else if p.StartedAt != nil {
			when = fmt.Sprintf("Started %s", humanize.Time(p.StartedAt.AsTime()))
		}
		rows = append(
			rows,
			table.Row{
				p.Pid,
				p.Configuration.Name,
				format.Truncate(processComandWithArguments(p), 30),
				when,
				prettyProcessStatus(p),
			},
		)
	}

	return format.NewTable(
		table.Row{"pid", "name", "command", "when", "status"},
		rows,
	).Render()
}

func (pto *ProcessTableOutput) Object() interface{} {
	if pto.processes == nil {
		return []*pb.ProcessInformation{}
	}
	return pto.processes
}
