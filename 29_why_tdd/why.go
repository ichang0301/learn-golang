// https://quii.gitbook.io/learn-go-with-tests/meta/why

// Lehman's law of software evolution: https://en.wikipedia.org/wiki/Lehman's_laws_of_software_evolution

package why

func Hello(name, language string) string {

	if language == "es" {
		return "Hola, " + name
	}

	if language == "fr" {
		return "Bonjour, " + name
	}

	// imagine dozens more languages

	return "Hello, " + name
}
