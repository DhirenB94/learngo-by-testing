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

type PlayerServer struct {
	Store PlayerStore
	http.Handler
}

//EMBEDDING
//We changed the second property of PlayerServer, removing the named property 'router http.ServeMux' and replaced it with http.Handler; this is called embedding.
//Go does not provide the typical, type-driven notion of subclassing, but it does have the ability to “borrow” pieces of an implementation by embedding types within a struct or interface.
//What this means is that our PlayerServer now has all the methods that http.Handler has, which is just ServeHTTP.
//To "fill in" the http.Handler we assign it to the router we create in NewPlayerServer.
//We can do this because http.ServeMux has the method ServeHTTP.
//This lets us remove our own ServeHTTP method, as we are already exposing one via the embedded type.

//You must be careful with embedding types because you will expose all public methods and fields of the type you embed.
//In our case, it is ok because we embedded just the interface that we wanted to expose (http.Handler).
//If we had been lazy and embedded http.ServeMux instead (the concrete type) it would still work
//but users of PlayerServer would be able to add new routes to our server because Handle(path, handler) would be public.

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
