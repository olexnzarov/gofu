package managedprocess

import (
	"time"

	"github.com/olexnzarov/gofu/pb"
	"google.golang.org/protobuf/types/known/durationpb"
)

const DEFAULT_RESTART_DELAY = time.Second * 10

type ProcessData struct {
	Id            string
	Configuration *pb.ProcessConfiguration
}

func (d *ProcessData) SetConfiguration(config *pb.ProcessConfiguration) {
	d.Configuration = config
}

func (d *ProcessData) GetConfiguration() *pb.ProcessConfiguration {
	return d.Configuration
}

func (d *ProcessData) GetID() string {
	return d.Id
}

func (d *ProcessData) GetRestartPolicy() *pb.ProcessConfiguration_RestartPolicy {
	if d.Configuration.RestartPolicy == nil {
		return &pb.ProcessConfiguration_RestartPolicy{
			AutoRestart: false,
			Delay:       durationpb.New(time.Duration(DEFAULT_RESTART_DELAY)),
			MaxRetries:  1,
		}
	}
	return d.Configuration.RestartPolicy
}

func (d *ProcessData) GetRestartDelay() time.Duration {
	policy := d.GetRestartPolicy()
	if policy.Delay == nil {
		return DEFAULT_RESTART_DELAY
	}
	return policy.Delay.AsDuration()
}
