package config

import (
	"os"

	"go1f/pkg/loger"

	"github.com/joho/godotenv"
)

var Cfg *Config

type Config struct {
	TODO_PORT string
}

func Init() {
	Cfg = &Config{}
	if err := godotenv.Load(); err != nil {
		loger.L.Error("Error loading .env file")
	}

	Cfg.TODO_PORT = os.Getenv("TODO_PORT")
	if Cfg.TODO_PORT == "" {
		Cfg.TODO_PORT = "7540"
	}
}
