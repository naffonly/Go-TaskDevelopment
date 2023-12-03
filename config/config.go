package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	KeyNotion   string
	Version     string
	DB_ID       string
	WebHook_URL string
	WebHook_ID  string
}

func InitConfig() *Config {
	var result = new(Config)
	result = loadConfig()

	if result == nil {
		log.Fatal("Config : Cannot start program, failed to load configuration")
		return nil
	}

	return result
}
func loadConfig() *Config {
	var result = new(Config)

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Config: Cannot load config file,", err.Error())
		return nil
	}

	if value, found := os.LookupEnv("TRANSTRACK_NOTION"); found {
		result.KeyNotion = value
	}
	if value, found := os.LookupEnv("VERSION"); found {
		result.Version = value
	}
	if value, found := os.LookupEnv("DB_ID"); found {
		result.DB_ID = value
	}
	if value, found := os.LookupEnv("WEBHOOK_URL"); found {
		result.WebHook_URL = value
	}
	if value, found := os.LookupEnv("WEBHOOK_ID"); found {
		result.WebHook_ID = value
	}

	return result
}
