package httpserver

import (
	"log"
	"net/http"

	httpserver "github.com/Skylli202/learn-go-with-tests/http-server/http-server"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemoryPlayerStore) RecordWin(name string) {}

func main() {
	server := &httpserver.PlayerServer{Store: &InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
