package main

import (
	httpserver "learngo/app/http-server"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(httpserver.PlayerServer)
	log.Fatal(http.ListenAndServe(":5001", handler))
}
