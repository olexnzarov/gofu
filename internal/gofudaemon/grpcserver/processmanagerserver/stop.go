package processmanagerserver

import (
	"context"

	"github.com/olexnzarov/gofu/pb"
)

func (s *Server) Stop(ctx context.Context, in *pb.StopRequest) (*pb.StopReply, error) {
	_, err := s.service.StopProcess(in.Process)
	if err != nil {
		return &pb.StopReply{
			Error: &pb.Error{
				Message: err.Error(),
			},
		}, nil
	}
	return &pb.StopReply{}, nil
}
