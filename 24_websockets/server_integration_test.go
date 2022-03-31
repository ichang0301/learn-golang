package command_line_time

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "[]")
	defer cleanDatabase()

	store, err := NewFileSystemPlayerStore(database)
	assertNoError(t, err)

	server := NewPlayerServer(store)
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

		var got League
		if err := json.NewDecoder(response.Body).Decode(&got); err != nil {
			t.Fatal(err)
		}

		assertStatus(t, response, http.StatusOK)
		assertLeague(t, got, League{{Name: "Pepper", Wins: 3}})
	})
}
