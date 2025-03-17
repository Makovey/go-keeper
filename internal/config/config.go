package config

type Config interface {
	DatabaseDSN() string
	GRPCPort() string
}

type config struct {
	databaseDSN string
	grpcPort    string
}

func (c *config) DatabaseDSN() string {
	return c.databaseDSN
}

func (c *config) GRPCPort() string {
	return c.grpcPort
}

func NewConfig() Config {
	cfg := newEnvConfig()

	return &config{
		databaseDSN: cfg.databaseDSN,
		grpcPort:    cfg.grpcPort,
	}
}
