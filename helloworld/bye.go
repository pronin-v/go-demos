package main

func Bye(name, lang string) string {
	if name == "" {
		return goodbyePrefix(lang) + goodbyeDefaultName(lang)
	}

	return goodbyePrefix(lang) + name
}

func goodbyePrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = "Adios, "
	default:
		prefix = "Bye bye, "
	}
	return
}

func goodbyeDefaultName(lang string) (name string) {
	switch lang {
	case spanish:
		name = "amigo"
	default:
		name = "Dude"
	}
	return
}
