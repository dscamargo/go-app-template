package pkg

import "os"

func GetEnvOrDefault(key string, defaultValue string) string {
	value, defined := os.LookupEnv(key)
	if !defined {
		return defaultValue
	}
	return value
}
