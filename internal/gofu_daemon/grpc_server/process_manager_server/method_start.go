package process_manager_server

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/olexnzarov/gofu/internal/gofu_daemon/process_manager"
	"github.com/olexnzarov/gofu/pb"
)

func (s *ProcessManagerServer) Start(ctx context.Context, in *pb.StartRequest) (*pb.StartReply, error) {
	if in.Configuration == nil {
		return nil, errors.New("configuration is expected, got nil")
	}

	processData := process_manager.ProcessData{
		Id:            uuid.New().String(),
		Configuration: in.Configuration,
	}
	s.NormalizeProcessData(&processData)
	process, err := s.processManager.Start(&processData)

	if err != nil {
		return &pb.StartReply{
			Response: &pb.StartReply_Error{
				Error: &pb.Error{Message: err.Error()},
			},
		}, nil
	}

	return &pb.StartReply{
		Response: &pb.StartReply_Process{
			Process: GetProcessInformation(process),
		},
	}, nil
}
