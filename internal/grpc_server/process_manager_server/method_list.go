package process_manager_server

import (
	"context"

	"github.com/alexnzarov/gofu/pb"
)

func (s *ProcessManagerServer) List(ctx context.Context, in *pb.ListRequest) (*pb.ListReply, error) {
	watchedProcesses := *s.processRegistry.Processes()
	processes := make([]*pb.ProcessInformation, 0, len(watchedProcesses))

	for _, p := range watchedProcesses {
		info := &pb.ProcessInformation{
			Id:            p.Data().Id,
			Pid:           int64(p.Pid()),
			Configuration: p.Data().Configuration,
			ExitState:     GetExitState(p),
			Status:        p.Status(),
		}
		processes = append(processes, info)
	}

	return &pb.ListReply{
		Processes: processes,
	}, nil
}
