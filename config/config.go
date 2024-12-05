package config

import "golang-fiber-api-template/internal/utils"

type Config struct {
	DbUrl           string
	DbDriverName    string
	DbMigrationPath string
	ApiUser         string
	ApiPassword     string
}

var Cfg *Config

func LoadConfig() {
	Cfg = &Config{
		DbUrl:           utils.GetFromEnv("DB_URL", "localhost"),
		DbDriverName:    utils.GetFromEnv("DB_DRIVER_NAME", "postgres"),
		DbMigrationPath: utils.GetFromEnv("DB_MIGRATION_PATH", "file://database/migrations"),
		ApiUser:         utils.GetFromEnv("API_USER", "admin"),
		ApiPassword:     utils.GetFromEnv("API_PASSWORD", "password"),
	}
}
