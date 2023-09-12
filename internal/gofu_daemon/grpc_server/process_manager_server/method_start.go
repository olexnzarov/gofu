package process_manager_server

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/olexnzarov/gofu/internal/gofu_daemon/process_manager"
	"github.com/olexnzarov/gofu/pb"
	"google.golang.org/protobuf/types/known/durationpb"
)

// SanitizeProcessName removes any illegal symbols from the given name.
// If the name will be still inadequate, returns an error.
func SanitizeProcessName(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return name, errors.New("process name is empty")
	}
	return name, nil
}

// GetDefaultProcessName returns a name for the process based on its command and arguments.
func GetDefaultProcessName(in *pb.ProcessConfiguration) string {
	return strings.Join(append([]string{in.Command}, in.Arguments...), " ")
}

func (s *ProcessManagerServer) Start(ctx context.Context, in *pb.StartRequest) (*pb.StartReply, error) {
	if in.Configuration == nil {
		return nil, errors.New("configuration is expected, got nil")
	}

	processData := process_manager.ProcessData{
		Id: uuid.New().String(),
	}

	// Sanitize the process name
	if sanitizedName, err := SanitizeProcessName(in.Configuration.Name); err == nil {
		in.Configuration.Name = sanitizedName
	} else {
		in.Configuration.Name = GetDefaultProcessName(in.Configuration)

	}

	// Set default restart policy
	if in.Configuration.RestartPolicy == nil {
		in.Configuration.RestartPolicy = &pb.ProcessConfiguration_RestartPolicy{
			AutoRestart: false,
			Delay:       durationpb.New(time.Duration(0)),
			MaxRetries:  1,
		}
	}

	processData.Configuration = in.Configuration

	mp, err := s.processManager.Start(&processData)

	if err != nil {
		return &pb.StartReply{
			Response: &pb.StartReply_Error{
				Error: &pb.Error{Message: err.Error()},
			},
		}, nil
	}

	return &pb.StartReply{
		Response: &pb.StartReply_Process{
			Process: &pb.ProcessInformation{
				Id:            processData.Id,
				Pid:           int64(mp.Pid()),
				Configuration: processData.Configuration,
				Status:        mp.Status(),
				ExitState:     GetExitState(mp),
			},
		},
	}, nil
}