package main

import (
	"log"
	"net/http"
)

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	storage map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.storage[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.storage[name]++
}

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
