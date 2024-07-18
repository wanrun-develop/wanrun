package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	postgresUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(postgresUrl), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected to DB")
	return db
}
