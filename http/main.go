package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	handler := &PlayerServer{&InMemoryPlayerStore{}}
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
