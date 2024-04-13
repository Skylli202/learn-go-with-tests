package httpserver

import (
	"io"
)

type FileSystemStore struct {
	Database io.Reader
}

func (f *FileSystemStore) GetLeague() []Player {
	league, _ := NewLeague(f.Database)

	return league
}
