package database

import (
	"github.com/golang-migrate/migrate/v4"
	"golang-fiber-api-template/config"
	"log"
)

func RunMigrations() {
	cfg := config.Cfg

	m, err := migrate.New(cfg.DbMigrationPath, cfg.DbUrl)

	if err != nil {
		log.Fatalf("Błąd migracji: %v", err)
	}

	if err := m.Up(); err != nil {
		log.Fatalf("Błąd migracji: %v", err)
	}

	log.Println("Migracje wykonane")
}
