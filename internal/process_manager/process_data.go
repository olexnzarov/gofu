package process_manager

import (
	"time"

	"github.com/olexnzarov/gofu/pb"
)

type ProcessData struct {
	Id            string
	Configuration *pb.ProcessConfiguration
}

const DEFAULT_RESTART_DELAY = time.Second * 10

func (d *ProcessData) AutoRestartDelay() time.Duration {
	if d.Configuration.RestartPolicy == nil || d.Configuration.RestartPolicy.Delay == nil {
		return DEFAULT_RESTART_DELAY
	}

	return d.Configuration.RestartPolicy.Delay.AsDuration()
}
