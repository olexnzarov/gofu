package procservice

import (
	"github.com/olexnzarov/gofu/internal/gofudaemon/procmanager"
	"github.com/olexnzarov/gofu/internal/logger"
)

type Storage interface {
	List() ([]*procmanager.ProcessData, error)
	Upsert(process *procmanager.ProcessData) error
	Delete(id string) error
}

type Service struct {
	log     logger.Logger
	manager *procmanager.Manager
	storage Storage
}

func New(log logger.Logger, manager *procmanager.Manager, storage Storage) *Service {
	return &Service{
		log:     log,
		manager: manager,
		storage: storage,
	}
}
