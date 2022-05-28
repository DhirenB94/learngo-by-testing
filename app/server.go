package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//We know the data is in our FakePlayerStore
//and we've abstracted that away into an interface PlayerStore.
//We need to update this so anyone passing us in a PlayerStore can provide us with the data for leagues.

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	Store PlayerStore
	http.Handler
}

type Player struct {
	Name string
	Wins int
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	ps := new(PlayerServer)

	ps.Store = store

	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(ps.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(ps.playerHandler))

	ps.Handler = router

	return ps
}

//Now we can update our handler code to call that rather than returning a hard-coded list.
//Delete our method getLeagueTable() and then update leagueHandler to call GetLeague().

func (ps PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(ps.Store.GetLeague())

	w.WriteHeader(http.StatusOK)
}

func (ps PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		ps.processWin(w, player)
	case http.MethodGet:
		ps.showScore(w, player)
	}
}

func (ps PlayerServer) showScore(w http.ResponseWriter, player string) {

	playersScore := ps.Store.GetPlayerScore(player)

	if playersScore == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, playersScore)
}

func (ps *PlayerServer) processWin(w http.ResponseWriter, player string) {

	ps.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
