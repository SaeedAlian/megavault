package main

import (
	"log"

	"megavault/api/cmd/api"
)

func main() {
	server := api.NewServer(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
