package process_registry

import (
	"time"

	"github.com/olexnzarov/gofu/pb"
)

type ProcessData struct {
	Id            string
	Configuration *pb.ProcessConfiguration
	Restarts      int
	autoRestarts  int
}

const DEFAULT_RESTART_DELAY = time.Second * 10

// HasAutoRestart returns whether the restart policy allows for this process to have an autorestart.
func (d *ProcessData) HasAutoRestart() bool {
	if d.Configuration.RestartPolicy == nil || !d.Configuration.RestartPolicy.AutoRestart {
		return false
	}
	return d.Configuration.RestartPolicy.MaxRestarts == 0 || uint32(d.autoRestarts) < d.Configuration.RestartPolicy.MaxRestarts
}

// AutoRestartDelay returns the delay the restart policy has. If it's not specified, returns a default delay.
func (d *ProcessData) AutoRestartDelay() time.Duration {
	if d.Configuration.RestartPolicy == nil || d.Configuration.RestartPolicy.Delay == nil {
		return DEFAULT_RESTART_DELAY
	}

	return d.Configuration.RestartPolicy.Delay.AsDuration()
}
