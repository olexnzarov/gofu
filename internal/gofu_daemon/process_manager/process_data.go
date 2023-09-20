package process_manager

import (
	"time"

	"github.com/olexnzarov/gofu/pb"
	"google.golang.org/protobuf/types/known/durationpb"
)

type ProcessData struct {
	Id            string
	Configuration *pb.ProcessConfiguration
}

const DEFAULT_RESTART_DELAY = time.Second * 10

func (d *ProcessData) RestartPolicy() *pb.ProcessConfiguration_RestartPolicy {
	if d.Configuration.RestartPolicy == nil {
		return &pb.ProcessConfiguration_RestartPolicy{
			AutoRestart: false,
			Delay:       durationpb.New(time.Duration(DEFAULT_RESTART_DELAY)),
			MaxRetries:  1,
		}
	}
	return d.Configuration.RestartPolicy
}

func (d *ProcessData) AutoRestartDelay() time.Duration {
	policy := d.RestartPolicy()
	if policy.Delay == nil {
		return DEFAULT_RESTART_DELAY
	}
	return policy.Delay.AsDuration()
}
