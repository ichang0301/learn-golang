package why

import "testing"

func TestHello(t *testing.T) {
	name := "test"

	t.Run("greet in Spanish", func(t *testing.T) {
		lan := "es"

		got := Hello(name, lan)
		want := "Hola, test"

		assertHello(t, got, want)
	})

	t.Run("greet in Franch", func(t *testing.T) {
		lan := "fr"

		got := Hello(name, lan)
		want := "Bonjour, test"

		assertHello(t, got, want)
	})

	t.Run("greet in English", func(t *testing.T) {
		got := Hello(name, "")
		want := "Hello, test"

		assertHello(t, got, want)
	})
}

func assertHello(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %s, but want %s", got, want)
	}
}
