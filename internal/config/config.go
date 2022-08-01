package config

import (
	"fmt"
	"os"
	"strconv"
)

var serviceVersion = "local"

const (
	port            = "HTTP_PORT"
	debugMode       = "DEBUG"
	postgresHost    = "POSTGRES_HOST"
	postgresPort    = "POSTGRES_PORT"
	postgresUser    = "POSTGRES_USER"
	postgresPass    = "POSTGRES_PASSWORD"
	postgresDB      = "POSTGRES_DATABASE"
	postgresTimeOUT = "POSTGRES_TIMEOUT"
	twitterID       = "TWITTER_ID"
	twitterSecret   = "TWITTER_SECRET"
	twitterTokenURL = "TWITTER_TOKEN_URL"
)

// Config contains the service configuration variables.
type Config struct {
	Port            string
	GRPCPort        string
	Debug           bool
	DbPostgresUrl   string
	DbTimeOUT       int
	TwitterID       string
	TwitterSecret   string
	TwitterTokenURL string
}

// New returns a structure with the service configuration variables.
func New() Config {
	postgresURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv(postgresUser),
		os.Getenv(postgresPass),
		os.Getenv(postgresHost),
		os.Getenv(postgresPort),
		os.Getenv(postgresDB))

	return Config{
		Port:            GetEnvString(port, "8080"),
		Debug:           GetEnvBool(debugMode, false),
		DbPostgresUrl:   postgresURL,
		DbTimeOUT:       GetEnvInt(postgresTimeOUT, 15),
		TwitterID:       GetEnvString(twitterID, ""),
		TwitterSecret:   GetEnvString(twitterSecret, ""),
		TwitterTokenURL: GetEnvString(twitterTokenURL, ""),
	}
}

// GetVersion returns the API version.
func GetVersion() string {
	return serviceVersion
}

// GetEnvString returns the value as a string of the environment
// variable that matches the key, if not found it returns the default value.
func GetEnvString(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return defaultValue
}

// GetEnvBool returns the value as boolean of the environment
// variable that matches the key, if not found it returns the default value.
func GetEnvBool(key string, defaultValue bool) bool {
	if val := os.Getenv(key); val != "" {
		bVal, err := strconv.ParseBool(val)
		if err != nil {
			return defaultValue
		}
		return bVal
	}

	return defaultValue
}

// GetEnvInt returns the value as integer of the environment
// variable that matches the key, if not found it returns the default value.
func GetEnvInt(key string, def int) int {
	val, err := strconv.Atoi(GetEnvString(key, fmt.Sprintf("%d", def)))
	if err != nil {
		return def
	}

	return val
}
