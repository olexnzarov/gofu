package format

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/olexnzarov/gofu/internal/gofu_daemon/process_manager"
	"github.com/olexnzarov/gofu/pb"
)

type PrettyProcess struct {
	process *pb.ProcessInformation
}

func NewPrettyProcess(process *pb.ProcessInformation) *PrettyProcess {
	return &PrettyProcess{process: process}
}

func (p *PrettyProcess) Status() string {
	color := text.FgRed
	switch p.process.Status {
	case process_manager.STATUS_RUNNING:
		color = text.FgGreen
	case process_manager.STATUS_RESTARTING, process_manager.STATUS_STOPPED:
		color = text.FgYellow
	}
	return color.Sprint(strings.ToUpper(p.process.Status))
}

func (p *PrettyProcess) CommandWithArguments() string {
	return strings.Join(append([]string{p.process.Configuration.Command}, p.process.Configuration.Arguments...), " ")
}

func PrintProcess(process *pb.ProcessInformation) {
	PrintProcesses([]*pb.ProcessInformation{process})
}

func PrintProcesses(processes []*pb.ProcessInformation) {
	if len(processes) == 0 {
		fmt.Println("Found no processes.")
		return
	}

	rows := make([]table.Row, 0, len(processes))
	for _, p := range processes {
		pretty := NewPrettyProcess(p)
		rows = append(
			rows,
			table.Row{
				p.Pid,
				p.Configuration.Name,
				Truncate(pretty.CommandWithArguments(), 25),
				pretty.Status(),
			},
		)
	}

	NewTable(
		table.Row{"pid", "name", "command", "status"},
		rows,
	).Render()
}
