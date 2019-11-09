package migration

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// DoMigrate - runs migrations on mysql config
func DoMigrate(migrationsFolder, user, password, host, port, dbname, direction string) {
	log.Printf("running migrations %s ...", direction)
	connString := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

	m, err := migrate.New(
		migrationsFolder,
		connString)

	if err != nil {
		log.Println("Migrations initialization failed!")
		log.Fatal(err)
	}

	if direction == "down" {
		if err := m.Down(); err != nil {
			if err.Error() != "no change" { // no change is not an error actually
				log.Println("Down migrations failed!")
				log.Fatal(err)
			}
		}
		log.Println("migrations ran successfully!")
		return
	}

	if err := m.Up(); err != nil {
		if err.Error() != "no change" { // no change is not an error actually
			log.Println("Up migrations failed!")
			log.Fatal(err)
		}
	}
	log.Println("migrations ran successfully!")
}
