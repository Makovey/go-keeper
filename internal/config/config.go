package config

import "github.com/Makovey/go-keeper/internal/logger"

type Config interface {
	DatabaseDSN() string
	GRPCPort() string
	SecretKey() string
}

type config struct {
	databaseDSN string
	grpcPort    string
	secretKey   string
}

func (c *config) DatabaseDSN() string {
	return c.databaseDSN
}

func (c *config) GRPCPort() string {
	return c.grpcPort
}

func (c *config) SecretKey() string {
	return c.secretKey
}

func NewConfig(log logger.Logger) Config {
	cfg := newEnvConfig()

	log.Debug("DatabaseDSN: " + cfg.databaseDSN)
	log.Debug("GRPCPort: " + cfg.grpcPort)

	return &config{
		databaseDSN: cfg.databaseDSN,
		grpcPort:    cfg.grpcPort,
		secretKey:   cfg.secretKey,
	}
}
