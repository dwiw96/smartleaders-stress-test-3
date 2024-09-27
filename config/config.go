package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	SERVER_PORT    string
	DB_USERNAME    string
	DB_PASSWORD    string
	DB_HOST        string
	DB_PORT        string
	DB_NAME        string
	REDIS_HOST     string
	REDIS_PASSWORD string
}

func GetEnvConfig() *EnvConfig {
	log.Println("<- getEnvConfig()")

	initEnvConfig()

	var resEnvConfig EnvConfig
	resEnvConfig.SERVER_PORT = os.Getenv("SERVER_PORT")
	resEnvConfig.DB_USERNAME = os.Getenv("DB_USERNAME")
	resEnvConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	resEnvConfig.DB_HOST = os.Getenv("DB_HOST")
	resEnvConfig.DB_PORT = os.Getenv("DB_PORT")
	resEnvConfig.DB_NAME = os.Getenv("DB_NAME")
	resEnvConfig.REDIS_HOST = os.Getenv("REDIS_HOST")
	resEnvConfig.REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")

	log.Println("-> getEnvConfig()")
	return &resEnvConfig
}

func initEnvConfig() {
	log.Println("<- initEnvConfig()")

	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	projectPath := filepath.Dir((basePath))
	envPath := filepath.Join(projectPath, ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Println("failed to load .env file, msg:", err)
		return
	}

	log.Println("-> initEnvConfig()")
}
