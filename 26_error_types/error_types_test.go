package error_types

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDumbGetter(t *testing.T) {
	t.Run("when you don't get a 200 you get a status error", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusTeapot)
		}))
		defer svr.Close()

		_, err := DumbGetter(svr.URL)
		if err == nil {
			t.Fatal("expected an error")
		}

		want := BadStatusError{URL: svr.URL, Status: http.StatusTeapot}
		got, isStatusErr := err.(BadStatusError)
		if !isStatusErr {
			t.Fatalf("was not a BadStatusError, got %T", err)
		}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
