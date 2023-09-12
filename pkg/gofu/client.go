package gofu

import (
	"github.com/olexnzarov/gofu/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	ProcessManager pb.ProcessManagerClient
}

func NewClient(config *Config) (*Client, error) {
	connection, err := grpc.Dial(config.DaemonTarget, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &Client{
		ProcessManager: pb.NewProcessManagerClient(connection),
	}, nil
}

// DefaultClient returns a client with the default configuration.
func DefaultClient() (*Client, error) {
	config, err := DefaultConfig()
	if err != nil {
		return nil, err
	}
	return NewClient(config)
}
