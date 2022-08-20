package main

import "fmt"

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, langage string) string {
	if name == "" {
		name = "World"
	}

	if langage == spanish {
		return spanishHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chils", ""))
}
