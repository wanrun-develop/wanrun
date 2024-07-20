package db

import (
	"fmt"
	"log"

	_ "github.com/joho/godotenv"
	"github.com/wanrun-develop/wanrun/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	config := configs.AppConfig

	postgresUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.PostgresUser(),
		config.PostgresPassword(),
		config.PostgresHost(),
		config.PostgresPort(),
		config.PostgresDB())

	db, err := gorm.Open(postgres.Open(postgresUrl), &gorm.Config{})

	if err != nil {
		return db, err
	}
	log.Println("Connected to DB")
	return db, nil
}

func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	if err := sqlDB.Close(); err != nil {
		return err
	}
	return nil
}
