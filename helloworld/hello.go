package helloworld

const (
	portuguese            = "Portuguese"
	french                = "French"
	englishHelloPrefix    = "Hello, "
	portugueseHelloPrefix = "Ola, "
	frenchHelloPrefix     = "Bonjour, "
)

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingTranslation(language) + name
}

func greetingTranslation(language string) (prefix string) {
	switch language {
	case portuguese:
		prefix = portugueseHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}
