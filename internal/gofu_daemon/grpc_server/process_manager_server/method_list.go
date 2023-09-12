package process_manager_server

import (
	"context"

	"github.com/olexnzarov/gofu/pb"
)

func (s *ProcessManagerServer) List(ctx context.Context, in *pb.ListRequest) (*pb.ListReply, error) {
	watchedProcesses := *s.processManager.Processes.All()
	processes := make([]*pb.ProcessInformation, 0, len(watchedProcesses))

	for _, p := range watchedProcesses {
		data := p.Data()
		info := &pb.ProcessInformation{
			Id:            data.Id,
			Pid:           int64(p.Pid()),
			Configuration: data.Configuration,
			ExitState:     GetExitState(p),
			Status:        p.Status(),
		}
		processes = append(processes, info)
	}

	return &pb.ListReply{
		Processes: processes,
	}, nil
}
