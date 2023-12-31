package processmanagerservice

import (
	"strings"
	"time"

	"github.com/lucasepe/codename"
	"github.com/olexnzarov/gofu/internal/daemon/processmanager"
	"github.com/olexnzarov/gofu/pb"
	"google.golang.org/protobuf/types/known/durationpb"
)

const (
	DefaultAutoRestart       = false
	DefaultRestartDelay      = time.Second * 1
	DefaultRestartMaxRetries = 0
)

func (s *Service) GetDefaultProcessName(config *pb.ProcessConfiguration) string {
	if rand, err := codename.DefaultRNG(); err == nil {
		for i := 0; i < 3; i++ {
			name := codename.Generate(rand, 0)
			if _, err := s.manager.Processes.Find(name); err != nil {
				return name
			}
		}
	}
	return strings.Join(append([]string{config.Command}, config.Arguments...), " ")
}

func (s *Service) ValidateProcessConfig(config *pb.ProcessConfiguration) error {
	if config == nil {
		return processmanager.ErrMissingProcessConfiguration
	}
	if config.Name == "" {
		return processmanager.ErrInvalidProcessName
	}
	if config.Command == "" {
		return processmanager.ErrInvalidProcessCommand
	}
	return nil
}

func (s *Service) SanitizeProcessConfig(config *pb.ProcessConfiguration) {
	if config == nil {
		return
	}
	config.Name = sanitizeProcessName(config.Name)
	if len(config.Name) == 0 {
		config.Name = s.GetDefaultProcessName(config)
	}
	config.RestartPolicy = sanitizeRestartPolicy(config.RestartPolicy)
}

func sanitizeProcessName(name string) string {
	return strings.TrimSpace(name)
}

func sanitizeRestartPolicy(policy *pb.ProcessConfiguration_RestartPolicy) *pb.ProcessConfiguration_RestartPolicy {
	if policy == nil {
		return &pb.ProcessConfiguration_RestartPolicy{
			AutoRestart: DefaultAutoRestart,
			Delay:       durationpb.New(DefaultRestartDelay),
			MaxRetries:  DefaultRestartMaxRetries,
		}
	}
	if policy.Delay == nil {
		policy.Delay = durationpb.New(DefaultRestartDelay)
	}
	return policy
}
