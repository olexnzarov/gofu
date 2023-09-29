package processmanagerserver

import (
	"github.com/olexnzarov/gofu/internal/gofudaemon/procmanager"
	"github.com/olexnzarov/gofu/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToExitState(p *procmanager.ManagedProcess) *pb.ProcessInformation_ExitState {
	code, exitedAt, err := p.GetExitState()
	if err != nil {
		return nil
	}
	return &pb.ProcessInformation_ExitState{
		Code:     int64(code),
		ExitedAt: timestamppb.New(exitedAt),
	}
}

func ToProcessInformation(process *procmanager.ManagedProcess) *pb.ProcessInformation {
	processData := process.GetData()
	info := &pb.ProcessInformation{
		Id:            processData.Id,
		Pid:           int64(process.GetProcessId()),
		Configuration: processData.Configuration,
		Status:        process.GetStatus(),
		ExitState:     ToExitState(process),
		Stdout:        process.GetStdoutPath(),
		Restarts:      process.GetRestarts(),
		StartedAt:     nil,
	}

	if inner, err := process.GetRunningProcess(); err == nil {
		info.StartedAt = timestamppb.New(inner.StartedAt())
	}

	return info
}

func ToProcessInformationArray(processes []*procmanager.ManagedProcess) []*pb.ProcessInformation {
	mapped := make([]*pb.ProcessInformation, 0, len(processes))
	for _, p := range processes {
		mapped = append(mapped, ToProcessInformation(p))
	}
	return mapped
}
