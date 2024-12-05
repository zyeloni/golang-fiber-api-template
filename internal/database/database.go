package database

import (
	"database/sql"
	"fmt"
	"golang-fiber-api-template/config"
)

var DB *sql.DB

func Connect() error {
	cfg := config.Cfg
	var err error
	DB, err = sql.Open(cfg.DbDriverName, cfg.DbUrl)

	if err != nil {
		return fmt.Errorf("błąd połączenia z bazą danych: %w", err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("nie udało się nawiązać połączenia z bazą: %w", err)
	}

	return nil
}
