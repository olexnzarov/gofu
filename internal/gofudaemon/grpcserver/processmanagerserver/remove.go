package processmanagerserver

import (
	"context"

	"github.com/olexnzarov/gofu/pb"
)

func (s *ProcessManagerServer) Remove(ctx context.Context, in *pb.RemoveRequest) (*pb.RemoveReply, error) {
	err := s.service.RemoveProcess(in.Process)
	if err != nil {
		return &pb.RemoveReply{
			Error: &pb.Error{
				Message: err.Error(),
			},
		}, nil
	}
	return &pb.RemoveReply{}, nil
}
