package main

import "fmt"

type contactinfo struct {
	email   string
	phoneno int
}
type person struct {
	firstName string
	lastName  string
	contact   contactinfo //dont have to specify field name,single contactinfo can declare field type and value
}

func main() {
	// manno := person{firstName: "munir", lastName: "muhammad"}

	// manno.lastName = "ayuba"
	// fmt.Println(manno)
	// fmt.Printf("%+v", manno)

	manno := person{firstName: "munir",
		lastName: "muhammad",
		contact: contactinfo{
			email:   "muhammadmuneermo@gmial.com",
			phoneno: 800213,
		},
	}
	manno.updateName("sanba")
	manno.print()
}

func (p *person) updateName(newFirstname string) {
	p.firstName = newFirstname

}

func (p person) print() {
	fmt.Printf("%+v", p)
}
