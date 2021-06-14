package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreetings() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreetings())
}

func (englishBot) getGreetings() string {
	//VERY custom logic for generating custom greeting
	return "Hi There!"
}

func (spanishBot) getGreetings() string {
	//very custom logic for generating custom greeting
	return "Hola amigo!"
}
