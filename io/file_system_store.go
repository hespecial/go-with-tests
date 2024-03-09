package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) (wins int) {
	for _, v := range f.GetLeague() {
		if name == v.Name {
			wins = v.Wins
			break
		}
	}
	return
}
