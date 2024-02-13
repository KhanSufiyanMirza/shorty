package env

import (
	"os"
	"strconv"
)

func GetEnv(env, defaultValue string) string {
	environment := os.Getenv(env)
	if environment == "" {
		return defaultValue
	}

	return environment
}

func parseEnvToInt(envName, defaultValue string) int {
	num, err := strconv.Atoi(GetEnv(envName, defaultValue))

	if err != nil {
		return 0
	}

	return num
}
