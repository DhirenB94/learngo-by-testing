package httpserver

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

//our server should not know the score, so we use our GetPlayerScorefunc into an interface

//for our PlayerServer to use the PlayerStore we need to change it into a struct

type PlayerServer struct {
	store PlayerStore
}

//Now to implement the Handler interface (with the serveHttp method), we add a method to our new struct with the existing handler code we had
//this method is the same as the method on the handler interface

func (ps *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, ps.store.GetPlayerScore(player))
}
