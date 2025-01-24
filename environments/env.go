package environments

import "os"

func getenv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func IsDevelopment() bool {
	return GetServer().Environment == "development"
}
