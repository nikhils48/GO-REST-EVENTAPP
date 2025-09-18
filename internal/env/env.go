package env

import (
	"os"
	"strconv"
)

func GetEnvString(key, defaultvalue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultvalue
}

func GetEnvInt(key string, defaultvalue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultvalue
}
