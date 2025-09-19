package db

import (
	"fmt"
	"log"
	"tutorial/config"
	"tutorial/service/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true}), &gorm.Config{})

	if err != nil {
		log.Println("Failed to connect to database. \n", err)
	}

	logger.Log.Info().Msg("✅ Connected to Database")

	DB = db
}
