package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/SaeedAlian/megavault/api/api"
	"github.com/SaeedAlian/megavault/api/config"
	"github.com/SaeedAlian/megavault/api/db"
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
