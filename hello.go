package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const russianHelloPrefix = "Привет, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	switch lang {
	case "Spanish":
		return spanishHelloPrefix + name

	case "French":
		return frenchHelloPrefix + name

	case "Russian":
		return russianHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
