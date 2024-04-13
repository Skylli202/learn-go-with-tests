package httpserver

import (
	"encoding/json"
	"io"
)

type FileSystemStore struct {
	Database io.Reader
}

func (f *FileSystemStore) GetLeague() []Player {
	var league []Player

	json.NewDecoder(f.Database).Decode(&league)

	return league
}
