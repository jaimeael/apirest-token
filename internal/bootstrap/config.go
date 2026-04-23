package bootstrap

import "os"

type Config struct {
    DBUrl     string
    JWTSecret string
    Port      string
}

func LoadConfig() *Config {
    return &Config{
        DBUrl:     getEnv("DB_URL", "postgres://user:pass@db:5432/tasks?sslmode=disable"),
        JWTSecret: getEnv("JWT_SECRET", "supersecret"),
        Port:      getEnv("PORT", "8080"),
    }
}

func getEnv(key, fallback string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return fallback
}