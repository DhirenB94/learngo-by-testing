package main

import (
	httpserver "learngo/app/http-server"
	"log"
	"net/http"
)

type InMemoryPlayerStore struct {
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemoryPlayerStore) RecordWin(name string) {

}
func main() {
	server := &httpserver.PlayerServer{Store: &InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5001", server))

}

//Even though our tests are passing we don't really have working software.
//running main doesnt use the software as intended because we havent implemented PlayerStore correctly
//by focusing on our handler we have identified the interface that we need
//We could start writing some tests around our InMemoryPlayerStore but it's only here temporarily (replace with db).
//SO we will write an integration test between our PlayerServer and InMemoryPlayerStore
//This allows us to be confident our application is working, without having to directly test InMemoryPlayerStore.
//Not only that, but when we get around to implementing PlayerStore with a database, we can test that implementation with the same integration test.
