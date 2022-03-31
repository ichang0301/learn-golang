package websockets_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	poker "github.com/ichang0301/learn-golang/24_websockets"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "[]")
	defer cleanDatabase()

	store, err := poker.NewFileSystemPlayerStore(database)
	assertNoError(t, err)

	server := mustMakePlayerServer(t, store)
	player := "Pepper"
	wantedCount := 3

	for i := 0; i < wantedCount; i++ {
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	}

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response, http.StatusOK)
		assertResponseBody(t, response.Body.String(), strconv.Itoa(3))
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())

		var got poker.League
		if err := json.NewDecoder(response.Body).Decode(&got); err != nil {
			t.Fatal(err)
		}

		assertStatus(t, response, http.StatusOK)
		assertLeague(t, got, poker.League{{Name: "Pepper", Wins: 3}})
	})
}
