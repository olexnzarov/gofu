package gofudaemon

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/olexnzarov/gofu/internal/gofudaemon/grpcserver"
	"github.com/olexnzarov/gofu/internal/logger"
	"github.com/olexnzarov/gofu/pkg/gofu"
	"go.uber.org/multierr"
)

type Daemon struct {
	log         logger.Logger
	server      *grpcserver.Server
	directories *gofu.Directories
}

func New(
	log logger.Logger,
	server *grpcserver.Server,
	directories *gofu.Directories,
) *Daemon {
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

func NewDatabase(directories *gofu.Directories) (*sql.DB, error) {
	return sql.Open(
		"sqlite3",
		fmt.Sprintf("%s/daemon.db", directories.DataDirectory),
	)
}

func (d *Daemon) Start() error {
	if err := d.server.Start(); err != nil {
		return err
	}
	if err := d.directories.CreateAll(); err != nil {
		return err
	}
	return d.announce()
}

func (d *Daemon) Stop() error {
	return multierr.Combine(
		d.server.Stop(),
		os.Remove(d.directories.DaemonTargetFile()),
	)
}
