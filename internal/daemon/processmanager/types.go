package processmanager

import (
	"time"

	"github.com/olexnzarov/gofu/pb"
)

type Storage interface {
	List() ([]ProcessData, error)
	Upsert(process ProcessData) error
	Delete(id string) error
	Initialize() error
}

type ProcessData interface {
	GetID() string
	SetConfiguration(*pb.ProcessConfiguration)
	GetConfiguration() *pb.ProcessConfiguration
	GetRestartPolicy() *pb.ProcessConfiguration_RestartPolicy
	GetRestartDelay() time.Duration
}
