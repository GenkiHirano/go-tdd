package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string, langage string) string {
	if name == "" {
		name = "World"
	}

	if langage == "Spanish" {
		return "Hola, " + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chils", ""))
}
