package main

import (
	httpserver "learngo/app/http-server"
	"log"
	"net/http"
)

//by changing the interface, main no longer compiles
//Error:
/*
./main.go:20:37: cannot use &InMemoryPlayerStore{} (type *InMemoryPlayerStore) as type httpserver.PlayerStore in field value:
*InMemoryPlayerStore does not implement httpserver.PlayerStore (missing RecordWin method)
*/

type InMemoryPlayerStore struct {
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &httpserver.PlayerServer{Store: &InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5001", server))

}
