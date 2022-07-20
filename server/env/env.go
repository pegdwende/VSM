package env

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var envLoaded = false

func GetEnvVariable(name string) string {

	loadEnv()

	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	value, _ := viper.Get(name).(string)

	return value
}

func GetRequiredEnvVariable(name string) string {

	loadEnv()

	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	value, ok := viper.Get(name).(string)

	if !ok {
		panic(fmt.Sprintf("Could not find environment variable value for %s environment variable", name))
	}

	return value
}

func loadEnv() {

	if envLoaded == false {
		// SetConfigFile explicitly defines the path, name and extension of the config file.
		// Viper will use this and not check any of the config paths.
		// .env - It will search for the .env file in the current directory
		viper.SetConfigFile("../.env")

		// Find and read the config file
		err := viper.ReadInConfig()

		if err != nil {
			log.Fatalf("Error while reading config file %s", err)
		}

		envLoaded = true
	}
}
