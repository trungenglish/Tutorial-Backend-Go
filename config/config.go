package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	
}

func InitConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Không tìm thấy file .env, dùng biến môi trường hệ thống.")
	}

	return &Config{
		Port:   getEnv("PORT", "8080"),
		
	}

}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}