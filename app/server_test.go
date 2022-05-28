package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

//Next, we'll want to extend our test so that we can control exactly what data we want back.
//We can update the test to assert that the league table contains some players that we will stub in our store.

//Update FakePlayerStore to let it store a league, which is just a slice of Player.
//We'll store our expected data in there.

type FakePlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *FakePlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *FakePlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *FakePlayerStore) GetLeague() []Player {
	return s.league
}

func TestGetPlayers(t *testing.T) {
	//create our store
	store := FakePlayerStore{scores: map[string]int{
		"Pedro": 20,
		"Floyd": 10,
	}}

	//create our playerserver
	server := NewPlayerServer(&store)

	t.Run("return Pedro's score", func(t *testing.T) {
		request := newGetScoreRequest("Pedro")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")

	})
	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")

	})
	t.Run("return 404 if player not found", func(t *testing.T) {
		request := newGetScoreRequest("Bilbo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestStoreWins(t *testing.T) {
	store := FakePlayerStore{
		map[string]int{},
		nil,
		nil}

	server := NewPlayerServer(&store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Pedro"
		request := newPostWinRequest("Pedro")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin, want %d", len(store.winCalls), 1)
		}
		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		}
	})

}

func TestLeague(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{
				Name: "Joao",
				Wins: 5,
			},
			{
				Name: "Jose",
				Wins: 7,
			},
			{
				Name: "Antonio",
				Wins: 9,
			},
		}

		store := FakePlayerStore{
			scores:   nil,
			winCalls: nil,
			league:   wantedLeague,
		}
		server := NewPlayerServer(&store)

		req := newLeagueRequest()
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		got := getLeagueFromResponse(t, resp.Body)
		assertStatus(t, resp.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
		assertContentType(t, resp, jsonContentType)
	})
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("status codes should be the same, got %d, want %d", got, want)
	}
}

func getLeagueFromResponse(t testing.TB, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return
}

func assertLeague(t testing.TB, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}
