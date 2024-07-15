package main

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	sqlFilePath := "file://migration_sql"
	postgresUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"))
	m, err := migrate.New(sqlFilePath, postgresUrl)
	if err != nil {
		log.Fatal(err)
	}

	direction := os.Getenv("MIGRATION_DIRECTION")

	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Printf("Create table!!!")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Printf("Delete table!!!")
	default:
		log.Fatalf("Unknown migration direction: %s", direction)
	}

}

func P(t interface{}) {
	fmt.Println(reflect.TypeOf(t))
}
