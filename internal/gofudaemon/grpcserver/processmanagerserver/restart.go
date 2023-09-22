package processmanagerserver

import (
	"context"

	"github.com/olexnzarov/gofu/pb"
)

func (s *ProcessManagerServer) Restart(ctx context.Context, in *pb.RestartRequest) (*pb.RestartReply, error) {
	_, err := s.service.RestartProcess(in.Process)
	if err != nil {
		return &pb.RestartReply{
			Error: &pb.Error{
				Message: err.Error(),
			},
		}, nil
	}
	return &pb.RestartReply{}, nil
}
