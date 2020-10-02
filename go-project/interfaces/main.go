package main

import "fmt"

type bot interface {
	getgreeting() string
}

type englishbot struct{}
type spanishbot struct{}

func main() {
	eb := englishbot{}
	sb := spanishbot{}

	printgreeting(eb)
	printgreeting(sb)

}

// BAD
// func printgreeting(eb englishbot) {
// 	fmt.Println(eb.getgreeting())
// }

// func printgreeting(sb spanishbot) { 
// 	fmt.Println(sb.getgreeting())
// }

func printgreeting(b bot) {
	fmt.Println(b.getgreeting())
}
func (eb englishbot) getgreeting() string {
	return "hi there"
}

func (sb spanishbot) getgreeting() string {
	return "hola"
}
