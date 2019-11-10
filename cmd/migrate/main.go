package main

import (
	"flag"

	"github.com/dpogorelovsky/go-basic-web-app/app/migration"
	"github.com/dpogorelovsky/go-basic-web-app/config"
)

// for cmd usage
func main() {
	var migrationDirection = flag.String("direction", "up", "migrations direction: up or down")
	flag.Parse()
	// loading local configuration
	config.LoadConfig()

	migration.DoMigrate(config.Get("DB_MIGRATIONS_FOLDER"),
		config.Get("DB_USER"),
		config.Get("DB_PASS"),
		config.Get("DB_HOST"),
		config.Get("DB_PORT"),
		config.Get("DB_NAME"),
		*migrationDirection)
}
