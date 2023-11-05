package managedprocess

import (
	"github.com/olexnzarov/gofu/internal/event"
	"github.com/olexnzarov/gofu/internal/process"
)

type ProcessExitEvent struct {
	Process  *ManagedProcess
	Exited   *process.Process
	ExitCode int
}

type ProcessSpawnEvent struct {
	Process *ManagedProcess
	Spawned *process.Process
	Error   error
}

type ProcessForceRestartEvent struct {
	Process    *ManagedProcess
	Spawned    *process.Process
	WasRunning bool
	Error      error
}

type Events struct {
	OnProcessExit         *event.Event[ProcessExitEvent]
	OnProcessSpawn        *event.Event[ProcessSpawnEvent]
	OnProcessForceRestart *event.Event[ProcessForceRestartEvent]
}

func newEvents() *Events {
	return &Events{
		OnProcessExit:         event.New[ProcessExitEvent](),
		OnProcessSpawn:        event.New[ProcessSpawnEvent](),
		OnProcessForceRestart: event.New[ProcessForceRestartEvent](),
	}
}

func (manager *Manager) subscribe() {
	manager.Event.OnProcessExit.Subscribe(manager.onProcessExit).Order(1000)
	manager.Event.OnProcessSpawn.Subscribe(manager.onProcessSpawn).Order(1000)
	manager.Event.OnProcessForceRestart.Subscribe(manager.onProcessForceRestart).Order(1000)
}
