package util

import "os"

func GetEnvOrDefault(env string, defaultValue string) string {
	if os.Getenv(env) != "" {
		return os.Getenv(env)
	}
	return defaultValue
}
