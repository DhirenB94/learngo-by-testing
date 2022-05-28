package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

//It's quite odd (and inefficient) to be setting up a router as a request comes in and then calling it.
//What we ideally want to do is have some kind of NewPlayerServer function which will take our dependencies and do the one-time setup of creating the router.
//Each request can then just use that one instance of the router

//PlayerServer now needs to store a router

type PlayerServer struct {
	Store  PlayerStore
	router *http.ServeMux
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	ps := &PlayerServer{
		store,
		http.NewServeMux(),
	}
	ps.router.Handle("/league", http.HandlerFunc(ps.leagueHandler))
	ps.router.Handle("/players/", http.HandlerFunc(ps.playerHandler))

	return ps
}

//We have moved the routing creation out of ServeHTTP and into our NewPlayerServer so this only has to be done once, not per request.

func (ps *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ps.router.ServeHTTP(w, r)
}

func (ps PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
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
