package httpserver_test

import (
	"strings"
	"testing"

	httpserver "github.com/Skylli202/learn-go-with-tests/http-server/http-server"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)

		store := httpserver.FileSystemStore{database}

		got := store.GetLeague()

		want := []httpserver.Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)
	})
}
