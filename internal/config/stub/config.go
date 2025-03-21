package stub

import "github.com/Makovey/go-keeper/internal/config"

type configStub struct {
	databaseDSN string
	grpcPort    string
	secretKey   string
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
