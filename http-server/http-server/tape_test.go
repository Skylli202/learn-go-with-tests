package httpserver_test

import (
	"io"
	"testing"

	httpserver "github.com/Skylli202/learn-go-with-tests/http-server/http-server"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	tape := &httpserver.Tape{File: file}

	tape.Write([]byte("abc"))

	file.Seek(0, io.SeekStart)
	newFileContent, _ := io.ReadAll(file)

	got := string(newFileContent)
	want := "abc"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
