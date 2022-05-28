package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
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

func (ps PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	//endpoint currently does not return a body, so it cannot be parsed into JSON, so :
	leagueTable := []Player{
		{
			Name: "Jimbo",
			Wins: 20,
		},
	}

	json.NewEncoder(w).Encode(leagueTable)

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
