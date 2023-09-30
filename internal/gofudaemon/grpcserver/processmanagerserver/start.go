package processmanagerserver

import (
	"context"

	"github.com/olexnzarov/gofu/pb"
)

func (s *Server) Start(ctx context.Context, in *pb.StartRequest) (*pb.StartReply, error) {
	process, err := s.service.CreateProcess(in.Configuration)
	if err != nil {
		return &pb.StartReply{
			Response: &pb.StartReply_Error{
				Error: &pb.Error{Message: err.Error()},
			},
		}, nil
	}
	return &pb.StartReply{
		Response: &pb.StartReply_Process{
			Process: ToProcessInformation(process),
		},
	}, nil
}
