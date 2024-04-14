package main

import (
	"log"
	"net/http"
	"os"

	httpserver "github.com/Skylli202/learn-go-with-tests/http-server/http-server"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store := &httpserver.FileSystemPlayerStore{Database: db}
	server := &httpserver.PlayerServer{Store: store}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
