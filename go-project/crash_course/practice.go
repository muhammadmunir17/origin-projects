package main

import (
	"fmt"
	"strconv"
)

type user struct {
	name, username, email string
	id                    int
}

func (u user) addto() string {
	return "age is now " + strconv.Itoa(u.id+1000)
}

func main() {

	user1 := user{"muhammad munir", "manno007", "muhammadmuneermo@gmail.com", 01}
	fmt.Println(user1.username)
	fmt.Println(user1.addto())

}
