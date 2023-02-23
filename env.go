package juniper

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var envLoaded bool

// GetEnv (an environment variable) by name This should be used over os.Getenv to ensure that the dotenv file has been loaded
func GetEnv(key string, fallback ...string) string {
	loadEnv()

	val := os.Getenv(key)

	if val == "" && len(fallback) > 0 {
		return fallback[0]
	}

	return val
}

// GetEnvInt (an environment variable) by name This should be used over os.Getenv to ensure that the dotenv file has been loaded
// vaule will be returned as an int
func GetEnvInt(key string, fallback ...int) int {
	loadEnv()

	val := os.Getenv(key)

	if val == "" && len(fallback) > 0 {
		return fallback[0]
	}

	parsed, err := strconv.Atoi(val)
	if err != nil {
		return fallback[0]
	}

	return parsed
}

// loadEnv from the .env file specifed in the DOT_ENV variable
func loadEnv() {
	if envLoaded {
		return
	}

	if os.Getenv("DOT_ENV") != "" {
		if err := godotenv.Load(os.Getenv("DOT_ENV")); err != nil {
			envLoaded = true
		}
	} else if err := godotenv.Load(); err != nil {
		envLoaded = true
	}

}
