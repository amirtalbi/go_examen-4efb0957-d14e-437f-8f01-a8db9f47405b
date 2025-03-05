package config

import (
    "os"
    "log"
)

type AppConfig struct {
    APIPrefix  string
    JWTSecret  string
    ServerPort string
}

var Config AppConfig

func LoadConfig() {
    Config = AppConfig{
        APIPrefix:  getEnv("API_PREFIX", "567088a9-6689-4e67-b5e5-ed40ad0a830c"),
        JWTSecret:  getEnv("JWT_SECRET", "your-secret-key"),
        ServerPort: getEnv("PORT", "8080"),
    }
    
    log.Println("Configuration loaded successfully")
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}