package main

import (
	httpserver "learngo/app/http-server"
	"log"
	"net/http"
)

func main() {
	server := &httpserver.PlayerServer{}
	log.Fatal(http.ListenAndServe(":5001", server))

	//I can change this to pass  my PlayerServer as the handler
	//Because it has a ServeHttp method - same as the hanlder interface so it satisfies
}
