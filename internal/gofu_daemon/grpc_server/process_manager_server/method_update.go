package process_manager_server

import (
	"context"

	"slices"

	"github.com/olexnzarov/gofu/pb"
)

func (s *ProcessManagerServer) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateReply, error) {
	process, err := s.processManager.Processes.Find(in.Process)
	if err != nil {
		return &pb.UpdateReply{
			Response: &pb.UpdateReply_Error{
				Error: &pb.Error{Message: err.Error()},
			},
		}, nil
	}

	in.UpdateMask.Normalize()
	if !in.UpdateMask.IsValid(in.Configuration) {
		return &pb.UpdateReply{
			Response: &pb.UpdateReply_Error{
				Error: &pb.Error{Message: "invalid update mask"},
			},
		}, nil
	}

	if slices.Contains(in.UpdateMask.Paths, "name") {
		name, err := SanitizeProcessName(in.Configuration.Name)
		if err != nil {
			return &pb.UpdateReply{
				Response: &pb.UpdateReply_Error{
					Error: &pb.Error{Message: err.Error()},
				},
			}, nil
		}
		in.Configuration.Name = name
	}

	err = s.processManager.UpdateConfiguration(process, in.Configuration, in.UpdateMask)
	if err != nil {
		return &pb.UpdateReply{
			Response: &pb.UpdateReply_Error{
				Error: &pb.Error{Message: err.Error()},
			},
		}, nil
	}

	return &pb.UpdateReply{
		Response: &pb.UpdateReply_Process{
			Process: GetProcessInformation(process),
		},
	}, nil
}
