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

func (f *FileSystemStore) GetPlayerScore(name string) int {
	var wins int
	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}
	return wins
}
