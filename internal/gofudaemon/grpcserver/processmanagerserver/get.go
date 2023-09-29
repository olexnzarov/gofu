package processmanagerserver

import (
	"context"

	"github.com/olexnzarov/gofu/pb"
)

func (s *ProcessManagerServer) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetReply, error) {
	process, err := s.service.FindProcess(in.Process)
	if err != nil {
		return &pb.GetReply{
			Response: &pb.GetReply_Error{
				Error: &pb.Error{Message: err.Error()},
			},
		}, nil
	}
	return &pb.GetReply{
		Response: &pb.GetReply_Process{
			Process: ToProcessInformation(process),
		},
	}, nil
}
