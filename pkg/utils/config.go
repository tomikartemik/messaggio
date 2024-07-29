package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Host        string
	Port        string
	Username    string
	Password    string
	DBName      string
	SSLMode     string
	KafkaBroker string
}

func GetConfig() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("error with loading data from env: %s", err.Error())
	}
	cfg := &Config{
		Host:        os.Getenv("DB_HOST"),
		Port:        os.Getenv("DB_PORT"),
		Username:    os.Getenv("DB_USERNAME"),
		Password:    os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
		SSLMode:     os.Getenv("DB_SSLMODE"),
		KafkaBroker: os.Getenv("KAFKA_BROKER"),
	}
	return cfg
}
