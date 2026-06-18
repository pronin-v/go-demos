package helloworld

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	russianHelloPrefix = "Привет, "
)

// Returns greeting message for name using specified language
// or english by default
func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix

	case french:
		prefix = frenchHelloPrefix

	case russian:
		prefix = russianHelloPrefix

	default:
		prefix = englishHelloPrefix
	}
	return
}
