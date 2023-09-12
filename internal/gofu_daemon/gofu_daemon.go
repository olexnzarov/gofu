package gofu_daemon

import (
	"os"

	"github.com/olexnzarov/gofu/internal/gofu_daemon/grpc_server"
	"github.com/olexnzarov/gofu/pkg/gofu"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

type Daemon struct {
	log         *zap.Logger
	server      *grpc_server.Server
	directories *gofu.Directories
}

func NewDaemon(log *zap.Logger, server *grpc_server.Server, directories *gofu.Directories) *Daemon {
	return &Daemon{
		log:         log,
		server:      server,
		directories: directories,
	}
}

func (d *Daemon) announce() error {
	return os.WriteFile(
		d.directories.DaemonTargetFile(),
		[]byte(d.server.Target()),
		0644,
	)
}

func (d *Daemon) Start() error {
	err := d.server.Start()
	if err != nil {
		return err
	}
	err = d.directories.CreateAll()
	if err != nil {
		return err
	}
	return d.announce()
}

func (d *Daemon) Stop() error {
	return multierr.Combine(
		d.server.Stop(),
		os.Remove(d.directories.DaemonTargetFile()),
		d.log.Sync(),
	)
}
