package httpserver

import (
	"fmt"
	"net/http"
	"strings"
)

//We need to update PlayerServer's idea of what a PlayerStore is by changing the interface
//if we're going to be able to call RecordWin

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	Store PlayerStore
}

func (ps *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		ps.processWin(w)
	case http.MethodGet:
		ps.showScore(w, r)
	}
}

func (ps PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	playersScore := ps.Store.GetPlayerScore(player)

	if playersScore == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, playersScore)
}

func (ps *PlayerServer) processWin(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
	ps.Store.RecordWin("Bob")
}
