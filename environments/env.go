package environments

import (
	"os"

	"github.com/assiljaby/go-ledger-backend/pkg/logger"
	"github.com/joho/godotenv"
)

// LoadEnv Loads env variables in development.
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		logger.Info("was not able to load .env, proceeding to load OS env")
	}
}

func getenv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func IsDevelopment() bool {
	return GetServer().Environment == "development"
}
