package httpserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlayers(t *testing.T) {
	t.Run("return Pedro's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pedro", nil)
		response := httptest.NewRecorder()
		//spy on what our handler writes to the ResponseWriter

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("returns Floyd's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "players/Floyd", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "10"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}
