package main

import (
	httpserver "learngo/app/http-server"
	"log"
	"net/http"
)

//we have not passed in a PlayerStore.
//hard coded for now because we arent storing any data

type InMemoryPlayerStore struct {
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &httpserver.PlayerServer{Store: &InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5001", server))

}
