package processmanagerserver

import (
	"context"

	"github.com/olexnzarov/gofu/internal/daemon/logger"
	processmanagerservice "github.com/olexnzarov/gofu/internal/daemon/processmanager/service"
	"github.com/olexnzarov/gofu/pb"
	"github.com/olexnzarov/gofu/pkg/gofu"
)

type Server struct {
	pb.UnimplementedProcessManagerServer
	log         logger.Logger
	directories gofu.Directories
	service     *processmanagerservice.Service
}

func New(
	log logger.Logger,
	directories gofu.Directories,
	service *processmanagerservice.Service,
) pb.ProcessManagerServer {
	return &Server{
		log:         log,
		directories: directories,
		service:     service,
	}
}

func (s *Server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetReply, error) {
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

func (s *Server) List(ctx context.Context, in *pb.ListRequest) (*pb.ListReply, error) {
	processes := s.service.ListProcesses(&processmanagerservice.ProcessListFilter{})
	return &pb.ListReply{
		Processes: ToProcessInformationArray(processes),
	}, nil
}

func (s *Server) Remove(ctx context.Context, in *pb.RemoveRequest) (*pb.RemoveReply, error) {
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

func (s *Server) Restart(ctx context.Context, in *pb.RestartRequest) (*pb.RestartReply, error) {
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
