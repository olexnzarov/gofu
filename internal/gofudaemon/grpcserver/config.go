package grpcserver

type Config struct {
	Port int
}

func NewConfig() *Config {
	return &Config{
		Port: 50051,
	}
}
