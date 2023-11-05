package processmanagerservice

import (
	"github.com/olexnzarov/gofu/internal/daemon/logger"
	"github.com/olexnzarov/gofu/internal/daemon/processmanager"
	"github.com/olexnzarov/gofu/internal/daemon/processmanager/service/managedprocess"
	"github.com/olexnzarov/gofu/pkg/gofu"
)

type Service struct {
	log     logger.Logger
	manager *managedprocess.Manager
	storage processmanager.Storage
}

func New(log logger.Logger, directories gofu.Directories, storage processmanager.Storage) *Service {
	return &Service{
		log:     log,
		manager: managedprocess.NewManager(log, directories),
		storage: storage,
	}
}
