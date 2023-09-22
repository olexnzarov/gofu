package processmanagerserver

import (
	"github.com/olexnzarov/gofu/internal/gofudaemon/procmanager"
	"github.com/olexnzarov/gofu/pb"
)

func ToExitState(p *procmanager.ManagedProcess) *pb.ProcessInformation_ExitState {
	code, err := p.ExitCode()
	if err != nil {
		return nil
	}
	return &pb.ProcessInformation_ExitState{
		Code: int64(code),
	}
}

func ToProcessInformation(process *procmanager.ManagedProcess) *pb.ProcessInformation {
	processData := process.Data()
	return &pb.ProcessInformation{
		Id:            processData.Id,
		Pid:           int64(process.Pid()),
		Configuration: processData.Configuration,
		Status:        process.Status(),
		ExitState:     ToExitState(process),
	}
}

func ToProcessInformationArray(processes []*procmanager.ManagedProcess) []*pb.ProcessInformation {
	mapped := make([]*pb.ProcessInformation, 0, len(processes))
	for _, p := range processes {
		mapped = append(mapped, ToProcessInformation(p))
	}
	return mapped
}
