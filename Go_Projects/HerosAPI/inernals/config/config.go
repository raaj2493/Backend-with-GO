package config


import "os"

type config struct {
	DBHost string
	DBPort string
	DBUser string
	DBPassword string 
	DBName string 
}

func Load () config {
	return  config{
		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
        DBUser:     getEnv("DB_USER", "postgres"),
        DBPassword: getEnv("DB_PASSWORD", "postgres"),
        DBName:     getEnv("DB_NAME", "superhero_db"),
	}
}

func getEnv(key string, fallback string) string {
    value := os.Getenv(key)
    if value == "" {
        return fallback
    }
    return value
}