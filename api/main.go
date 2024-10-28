package main

import (
	"database/sql"
	"fmt"
	"log"

	"megavault/api/api"
	"megavault/api/config"
	"megavault/api/db"
)

func main() {
	db, err := db.NewPGSQLStorage()
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewServer(fmt.Sprintf(":%s", config.Env.Port), db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connection to DB was successful.")
}
