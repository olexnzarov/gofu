package processmanagerserver

import (
	"context"

	"github.com/olexnzarov/gofu/internal/gofudaemon/procservice"
	"github.com/olexnzarov/gofu/pb"
)

func (s *ProcessManagerServer) List(ctx context.Context, in *pb.ListRequest) (*pb.ListReply, error) {
	processes := s.service.ListProcesses(&procservice.ProcessListFilter{})
	return &pb.ListReply{
		Processes: ToProcessInformationArray(processes),
	}, nil
}
