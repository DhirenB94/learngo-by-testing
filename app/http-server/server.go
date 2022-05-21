package httpserver

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	Store PlayerStore
}

func (ps *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	playersScore := ps.Store.GetPlayerScore(player)

	if playersScore == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, playersScore)
}
