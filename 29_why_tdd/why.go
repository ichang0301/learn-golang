// https://quii.gitbook.io/learn-go-with-tests/meta/why

// Lehman's law of software evolution: https://en.wikipedia.org/wiki/Lehman's_laws_of_software_evolution

package why

import "fmt"

func Hello(name, language string) string {
	return fmt.Sprintf(
		"%s, %s",
		greeting(language),
		name,
	)
}

var greetings = map[string]string{
	"es": "Hola",
	"fr": "Bonjour",
}

func greeting(language string) string {
	greeting, exists := greetings[language]

	if exists {
		return greeting
	}

	return "Hello"
}
