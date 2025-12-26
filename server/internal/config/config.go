package config

import "os"

type Config struct {
	Server ServerConfig
	Mongo  MongoConfig
}

type ServerConfig struct {
	Port string
}

type MongoConfig struct {
	URI      string
	Database string
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
		Mongo: MongoConfig{
			URI:      getEnv("MONGO_URI", "mongodb://localhost:27017"),
			Database: getEnv("MONGO_DB", "app_db"),
		},
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
