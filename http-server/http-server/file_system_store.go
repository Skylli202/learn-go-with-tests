package httpserver

import (
	"encoding/json"
	"io"
)

type FileSystemStore struct {
	Database io.ReadWriteSeeker
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

func (f *FileSystemStore) RecordWin(name string) {
	league := f.GetLeague()
	for i, player := range league {
		if player.Name == name {
			league[i].Wins += 1
		}
	}
	f.Database.Seek(0, io.SeekStart)
	json.NewEncoder(f.Database).Encode(league)
}
