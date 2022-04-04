package error_types

import (
	"errors"
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
		var got BadStatusError

		isBadStatusError := errors.As(err, &got) // errors.As to try and extract our error into our custom type. It returns a bool to denote success and extracts it into got for us. go blog about errors.Is and errors.As (>= go 1.13) : https://go.dev/blog/go1.13-errors
		if !isBadStatusError {
			t.Fatalf("was not a BadStatusError, got %T", err)
		}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
