package hello

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want) // "%q" is very useful as it wraps your values in double quotes.
		}
	}

	t.Run("to a person", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", spanish)
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Lauren", french)
		want := "Bonjour, Lauren"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Korean", func(t *testing.T) {
		got := Hello("Lauren", korean)
		want := "안녕, Lauren"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("Lauren", japanese)
		want := "こんにちは、Lauren"
		assertCorrectMessage(t, got, want)
	})
}
