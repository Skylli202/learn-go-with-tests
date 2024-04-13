package httpserver

import (
	"io"
)

type FileSystemStore struct {
	Database io.ReadSeeker
}

func (f *FileSystemStore) GetLeague() []Player {
	f.Database.Seek(0, io.SeekStart)
	league, _ := NewLeague(f.Database)

	return league
}
