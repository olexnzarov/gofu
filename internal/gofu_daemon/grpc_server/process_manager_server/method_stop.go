package process_manager_server

import (
	"context"

	"github.com/olexnzarov/gofu/pb"
)

func (s *ProcessManagerServer) Stop(ctx context.Context, in *pb.StopRequest) (*pb.StopReply, error) {
	process, err := s.processManager.Processes.Find(in.Process)
	if err != nil {
		return &pb.StopReply{
			Error: &pb.Error{
				Message: err.Error(),
			},
		}, nil
	}

	err = process.Stop()

	if err != nil {
		return &pb.StopReply{
			Error: &pb.Error{
				Message: err.Error(),
			},
		}, nil
	}

	return &pb.StopReply{}, nil
}
