package httpserver

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

//tests were failing because we weren't passing anything into our playerserver (needs a store)
//so we create a fake store
type FakePlayerStore struct {
	scores map[string]int
}

func (s *FakePlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
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

		assertResponseBody(t, response.Body.String(), "20")
	})
	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "10")

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
