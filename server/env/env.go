package env

import (
	"fmt"
	"os"
)

func GetEnvVariable(name string, required bool) string {
	return os.Getenv(name)
}

func GetRequiredEnvVariable(name string) string {

	envValue := os.Getenv(name)

	if len(envValue) <= 0 {
		panic(fmt.Sprintf("Could not find environment variable value for %s environment variable", name))
	}

	return envValue
}
