package grpc

type ServerConfig struct {
	Port int
}

func DefaultServerConfig() ServerConfig {
	return ServerConfig{
		Port: 50051,
	}
}
