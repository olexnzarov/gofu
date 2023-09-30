package processmanagerserver

import (
	"context"

	"github.com/olexnzarov/gofu/pb"
)

func (s *Server) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateReply, error) {
	process, err := s.service.UpdateProcess(in.Process, in.Configuration, in.UpdateMask)
	if err != nil {
		return &pb.UpdateReply{
			Response: &pb.UpdateReply_Error{
				Error: &pb.Error{Message: err.Error()},
			},
		}, nil
	}
	return &pb.UpdateReply{
		Response: &pb.UpdateReply_Process{
			Process: ToProcessInformation(process),
		},
	}, nil
}
