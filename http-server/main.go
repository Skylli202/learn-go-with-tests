package main

import (
	"log"
	"net/http"

	httpserver "github.com/Skylli202/learn-go-with-tests/http-server/http-server"
)

func main() {
	server := &httpserver.PlayerServer{Store: httpserver.NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
