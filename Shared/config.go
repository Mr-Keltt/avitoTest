package shared

import (
	"log"
	"os"
)

type Config struct {
	ServerAddress   string
	PostgresConn    string
	PostgresJDBCURL string
	PostgresUser    string
	PostgresPass    string
	PostgresHost    string
	PostgresPort    string
	PostgresDB      string
}

func LoadConfig() *Config {
	config := &Config{
		ServerAddress:   getEnv("SERVER_ADDRESS"),
		PostgresConn:    getEnv("POSTGRES_CONN"),
		PostgresJDBCURL: getEnv("POSTGRES_JDBC_URL"),
		PostgresUser:    getEnv("POSTGRES_USERNAME"),
		PostgresPass:    getEnv("POSTGRES_PASSWORD"),
		PostgresHost:    getEnv("POSTGRES_HOST"),
		PostgresPort:    getEnv("POSTGRES_PORT"),
		PostgresDB:      getEnv("POSTGRES_DATABASE"),
	}

	return config
}

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}
