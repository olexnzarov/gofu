package process_manager_server

import (
	"context"

	"github.com/olexnzarov/gofu/pb"
)

func (s *ProcessManagerServer) Remove(ctx context.Context, in *pb.RemoveRequest) (*pb.RemoveReply, error) {
	process, err := s.processManager.Processes.Find(in.Process)
	if err != nil {
		return &pb.RemoveReply{
			Error: &pb.Error{
				Message: err.Error(),
			},
		}, nil
	}

	err = s.processManager.Remove(process)
	if err != nil {
		return &pb.RemoveReply{
			Error: &pb.Error{
				Message: err.Error(),
			},
		}, nil
	}

	return &pb.RemoveReply{}, nil
}
