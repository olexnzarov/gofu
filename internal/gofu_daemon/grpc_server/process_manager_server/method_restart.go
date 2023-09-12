package process_manager_server

import (
	"context"

	"github.com/olexnzarov/gofu/pb"
)

func (s *ProcessManagerServer) Restart(ctx context.Context, in *pb.RestartRequest) (*pb.RestartReply, error) {
	process, err := s.processManager.Processes.Find(in.Process)
	if err != nil {
		return &pb.RestartReply{
			Error: &pb.Error{
				Message: err.Error(),
			},
		}, nil
	}

	err = process.Restart()

	if err != nil {
		return &pb.RestartReply{
			Error: &pb.Error{
				Message: err.Error(),
			},
		}, nil
	}

	return &pb.RestartReply{}, nil
}
