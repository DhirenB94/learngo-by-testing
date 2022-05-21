package httpserver

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

//Now we want to check that when we make the Post /players/{name},
//that our PlayerStore somehow records the win

type FakePlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *FakePlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *FakePlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGetPlayers(t *testing.T) {
	//create our store
	store := FakePlayerStore{scores: map[string]int{
		"Pedro": 20,
		"Floyd": 10,
	}}

	//create our playerserver
	server := &PlayerServer{&store}

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
		nil}

	server := &PlayerServer{Store: &store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Pedro", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin, want %d", len(store.winCalls), 1)
		}

	})

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
