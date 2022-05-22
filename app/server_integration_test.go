package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	//creating the 2 components we are trying to integrate with
	store := InMemoryPlayerStore{}
	server := PlayerServer{Store: &store}

	player := "Pedro"

	//We then fire off 3 requests to record 3 wins for player.
	//We're not too concerned about the status codes in this test as it's not relevant to whether they are integrating well.

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	//The next response we do care about (so we store a variable response)
	//because we are going to try and get the player's score.

	response := httptest.NewRecorder()

	server.ServeHTTP(response, newGetScoreRequest(player))

	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")

}
