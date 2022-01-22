package config

import (
	"os"

	log "github.com/robpaul9/golog"
)

type Config struct {
	ServiceName     string
	ServicePort     string
	DBName          string
	DBHost          string
	DBPort          string
	DBUser          string
	DBPassword      string
	DBPasswordParam string
	Logger          log.Logger
}

func NewConfig() *Config {
	return &Config{
		ServiceName:     os.Getenv("SERVICE_NAME"),
		ServicePort:     os.Getenv("SERVICE_PORT"),
		DBName:          os.Getenv("DATABASE_NAME"),
		DBHost:          os.Getenv("DATABASE_HOST"),
		DBPassword:      os.Getenv("DATABASE_PASSWORD"),
		DBPasswordParam: os.Getenv("DB_PASSWORD_PARAM"),
		DBUser:          os.Getenv("DATABASE_USER"),
		DBPort:          os.Getenv("DATABASE_PORT"),
		Logger:          log.New(log.Config{ServiceName: os.Getenv("SERVICE_NAME")}),
	}
}
