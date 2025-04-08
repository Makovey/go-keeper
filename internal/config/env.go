package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envConfig struct {
	databaseDSN          string
	grpcPort             string
	secretKey            string
	clientConnectionHost string
	updateUIDuration     string
}

func newEnvConfig() envConfig {
	fn := "config.newEnvConfig"

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("[%s]: %s", fn, err.Error())
	}

	return envConfig{
		databaseDSN:          os.Getenv("DATABASE_DSN"),
		grpcPort:             os.Getenv("GRPC_PORT"),
		secretKey:            os.Getenv("SECRET_KEY"),
		clientConnectionHost: os.Getenv("CLIENT_CONNECTION_HOST"),
		updateUIDuration:     os.Getenv("UPDATE_UI_DURATION"),
	}
}
