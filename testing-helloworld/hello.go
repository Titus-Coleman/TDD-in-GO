package main

import "fmt"

const englishGreeting = "Hello, "
const spanishGreeting = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	switch language {
	case "Spanish":
		return spanishGreeting + name
	default:
		return englishGreeting + name
	}
}

func main() {
	fmt.Println(Hello("", ""))
}
