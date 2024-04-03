package httpserver_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	httpsrv "github.com/Skylli202/learn-go-with-tests/http-server"
	httpserver "github.com/Skylli202/learn-go-with-tests/http-server/http-server"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := httpsrv.InMemoryPlayerStore{}
	server := httpserver.PlayerServer{&store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}
