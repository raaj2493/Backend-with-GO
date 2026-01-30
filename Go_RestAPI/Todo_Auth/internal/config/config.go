package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
    PORT string
    DATABASEURL string
}

func Load()(*Config, error){
    var err error = godotenv.Load()

    if err != nil {
        log.Println("Error while loading env file", err)
        return nil, err
    }

    var Config *Config = &Config{
        PORT: os.Getenv("PORT"),
        DATABASEURL: os.Getenv("DATABASE_URL"),
    }

	return Config, nil
}