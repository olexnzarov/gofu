package process_manager_server

import (
	"errors"
	"strings"

	"github.com/lucasepe/codename"
	"github.com/olexnzarov/gofu/internal/gofu_daemon/process_manager"
	"github.com/olexnzarov/gofu/pb"
)

func (s *ProcessManagerServer) NormalizeProcessData(data *process_manager.ProcessData) {
	if sanitizedName, err := SanitizeProcessName(data.Configuration.Name); err == nil {
		data.Configuration.Name = sanitizedName
	} else {
		data.Configuration.Name = s.GetDefaultProcessName(data.Configuration)
	}
	if data.Configuration.RestartPolicy == nil {
		data.Configuration.RestartPolicy = data.RestartPolicy()
	}
}

// GetExitState returns an exit state if process has exited, otherwise it returns nil.
func GetExitState(p *process_manager.ManagedProcess) *pb.ProcessInformation_ExitState {
	code, err := p.ExitCode()
	if err != nil {
		return nil
	}
	return &pb.ProcessInformation_ExitState{
		Code: int64(code),
	}
}

// SanitizeProcessName removes any illegal symbols from the given name.
// If the name will be still inadequate, returns an error.
func SanitizeProcessName(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return name, errors.New("process name is empty")
	}
	return name, nil
}

// GetDefaultProcessName returns a human-readable name for the process.
// If it can't generate a unique one, returns a name based on process' command and arguments.
func (s *ProcessManagerServer) GetDefaultProcessName(in *pb.ProcessConfiguration) string {
	if rand, err := codename.DefaultRNG(); err == nil {
		for i := 0; i < 3; i++ {
			name := codename.Generate(rand, 0)
			if _, err := s.processManager.Processes.Find(name); err != nil {
				return name
			}
		}
	}
	return strings.Join(append([]string{in.Command}, in.Arguments...), " ")
}

func GetProcessInformation(process *process_manager.ManagedProcess) *pb.ProcessInformation {
	processData := process.Data()
	return &pb.ProcessInformation{
		Id:            processData.Id,
		Pid:           int64(process.Pid()),
		Configuration: processData.Configuration,
		Status:        process.Status(),
		ExitState:     GetExitState(process),
	}
}
