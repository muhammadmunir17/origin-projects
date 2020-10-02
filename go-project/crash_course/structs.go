package main

import (
	"fmt"
	"strconv"
)

type student struct {
	// firstname string
	// lastname  string	 
	// regno     string
	firstname, lastname, regno string
	age                        int
}

//value reciever
func (s student) greet() string {
	return "hello, my name is" + s.firstname + " " + s.lastname + " " + "and i am " + strconv.Itoa(s.age)
}

// pointer reciever
func (s *student) newage() {
	s.age++
}

func main() {
	//long
	//student{firstname: "muhammad munir", lastname: "ayuba", regno: "u16ee1033", age: "21"}
	student1 := student{"muhammad munir", "ayuba", "u16ee1033", 21}

	student1.newage()
	fmt.Println(student1)
	fmt.Println(student1.greet())
}
