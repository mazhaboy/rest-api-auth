package main

import (
	server "github.com/rest-api-auth"
	"log"
)

func main() {
	srv := new(server.Server)
	err := srv.Run("8000")
	if err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
