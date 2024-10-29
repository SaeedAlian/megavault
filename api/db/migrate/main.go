package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	pgMigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"megavault/api/db"
)

func main() {
	db, err := db.NewPGSQLStorage()
	if err != nil {
		log.Fatal(err)
	}

	driver, err := pgMigrate.WithInstance(db, &pgMigrate.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrate/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	ver, dirty, _ := m.Version()
	log.Printf("Migration Version: %d, Is Dirty: %v", ver, dirty)

	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}
