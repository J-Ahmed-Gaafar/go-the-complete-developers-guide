package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//func printGreeting(eb englishBot) {
//	fmt.Println(eb.getGreeting())
//}

//func printGreeting(sb spanishBot) {
//	fmt.Println(sb.getGreeting())
//}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for generating a spanish greeting
	return "Hola!"
}
