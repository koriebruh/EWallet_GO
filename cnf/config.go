package cnf

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

type Config struct {
	Server   string
	DataBase string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		slog.Error("Warning can't find .env")
	}

	cnf := Config{
		Server:   os.Getenv("SERVER"),
		DataBase: os.Getenv("DATABASE"),
	}

	if cnf.DataBase == "" || cnf.Server == "" {
		slog.Error("Missing required .env variables")
	}

	return &cnf, nil
}
