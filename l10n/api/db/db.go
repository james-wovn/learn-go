package db

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("user=%s password=%s dbname=%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_DBNAME"),
		),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("failed to open db connection")
	}

	DB = db
}

func Migrate() {
	m, err := migrate.New(
		"file://db/migrations",
		fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_DBNAME"),
		))
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
