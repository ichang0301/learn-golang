// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world

package hello

const spanish = "Spanish"
const french = "French"
const korean = "Korean"
const japanese = "Japanese"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const koreanHelloPrefix = "안녕, "
const japaneseHelloPrefix = "こんにちは、"

// Hello returns a personalised greeting in a given language.
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) { // "prefix" is a named return value: https://github.com/golang/go/wiki/CodeReviewComments#named-result-parameters
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case korean:
		prefix = koreanHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
