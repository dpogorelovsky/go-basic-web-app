package mysql

import (
	"context"
	"fmt"
	"log"

	"github.com/dpogorelovsky/go-basic-web-app/app/storage/astorage"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// storage - its instance will have all implementation
type storage struct {
	driver *sqlx.DB
}

// NewStorage - storage constructor
// has to implement all abstract methods
func NewStorage(host, user, password, port, dbname string) (astorage.Storage, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
	db, err := sqlx.Connect("mysql", connString)
	if err != nil {
		log.Println("Failed to prepare DB connection")
		return nil, err
	}

	ctx := context.Background()
	if err := db.PingContext(ctx); err != nil {
		log.Println("Failed to connect DB")
		return nil, err
	}

	s := &storage{}
	s.driver = db

	return s, nil
}
