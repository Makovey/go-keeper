package config

import (
	"time"

	"github.com/Makovey/go-keeper/internal/logger"
)

type Config interface {
	DatabaseDSN() string
	GRPCPort() string
	SecretKey() string
	ClientConnectionHost() string
	UpdateDurationForUI() time.Duration
}

type config struct {
	databaseDSN          string
	grpcPort             string
	secretKey            string
	clientConnectionHost string
	durationForUpdate    string
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

func (c *config) ClientConnectionHost() string {
	return c.clientConnectionHost
}

func (c *config) UpdateDurationForUI() time.Duration {
	d, err := time.ParseDuration(c.durationForUpdate)
	if err != nil {
		return 30 * time.Second
	}

	return d
}

func NewConfig(log logger.Logger) Config {
	cfg := newEnvConfig()

	log.Debug("DatabaseDSN: " + cfg.databaseDSN)
	log.Debug("GRPCPort: " + cfg.grpcPort)
	log.Debug("ClientConnectionHost: " + cfg.clientConnectionHost)
	log.Debug("UpdateDurationForUI: " + cfg.updateUIDuration)

	return &config{
		databaseDSN:          cfg.databaseDSN,
		grpcPort:             cfg.grpcPort,
		secretKey:            cfg.secretKey,
		clientConnectionHost: cfg.clientConnectionHost,
		durationForUpdate:    cfg.updateUIDuration,
	}
}
