package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{Store: &InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5001", server))

}
