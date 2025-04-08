package stub

import (
	"time"

	"github.com/Makovey/go-keeper/internal/config"
)

type configStub struct {
	databaseDSN      string
	grpcPort         string
	secretKey        string
	clientHost       string
	updateUIDuration time.Duration
}

func NewStubConfig() config.Config {
	return &configStub{}
}

func (c *configStub) DatabaseDSN() string {
	return c.databaseDSN
}

func (c *configStub) GRPCPort() string {
	return c.grpcPort
}

func (c *configStub) SecretKey() string {
	return c.secretKey
}

func (c *configStub) ClientConnectionHost() string {
	return c.clientHost
}

func (c *configStub) UpdateDurationForUI() time.Duration {
	return c.updateUIDuration
}
