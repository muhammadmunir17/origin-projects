package main

import "fmt"

//functions
//func greeting(name string) string {
//return "welcome" + " " + name
//}

//func getsum(num1, num2 int) int {
//return num1 + num2
//}
func main() {

	//	var fruitarr [3]string

	//fruitarr[0] = "apple"
	//fruitarr[1] = "orange"
	//fruitarr[2] = "pineapple"

	// shorthand
	//fruitarr := [3]string{"pineapple", "orange", "apple"}

	//slices

	//fruitslice := []string{"pineapple", "orange", "grapes"}

	//fmt.Println(fruitslice[0:2])
	//fmt.Println(len(fruitslice[1]))

	//fmt.Println(fruitarr[2])
	//fmt.Println(greeting("munir"))
	//fmt.Println(getsum(3, 2))

	//CONDITIONALS
	//if else
	x := 31
	y := 40000
	name := "manno"

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d\n", x, y)
	}
	//else if
	if name == "manno" {
		fmt.Printf("manno is your nickaname\n")

	} else if name == "muchacho" {
		fmt.Printf("muchacho is your other nickname\n")

	} else {
		fmt.Printf("thats not ur name\n")
	}

	switch name {
	case "manno":
		fmt.Printf("\nmanno is your other nickname")

	case "muchacho":
		fmt.Printf("muchacho is your other nickname\n")

	default:
		fmt.Printf("not ur nickname\n")

	}

	//loops
	//long
	i := 1
	for i <= 5 {
		fmt.Println(i)
		//i = i + 1
		i++
	}

	//short

	for i := 1; i <= 10; i++ {
		fmt.Printf("you are number %d\n", i)
	}

	//fizzbuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}

	//MAPS
	//declare map
	//emails := make(map[string]string)

	// assign key values
	// emails["user1"] = "muhammadmuneeermo@gmail.com"
	// emails["user2"] = "muhammadmuneermo2@gmail.com"

	//declare and asssign
	emails := map[string]string{"user1": "muhammadmuneermo2gmail.com", "user2": "muhammadmuneermo2@gmail.com"}
	emails["new user"] = "lailahayyub@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["user1"])

	// delete(emails, "user1")
	// fmt.Println(emails)

	//range
	regno := []int{1003, 1013, 1023, 1033, 1043, 1053, 1063}

	//looping with index
	//for i, reg := range regno {
	// 	fmt.Printf("%d - REG.NO: %d\n", i, reg)
	// }

	//without index
	for _, reg := range regno {
		fmt.Printf("REG.NO: %d\n", reg)
	}

	sum := 0
	for _, reg := range regno {
		sum += reg
	}
	fmt.Println("Sum", sum)

	//range with maps

	// for k, v := range emails {
	// 	fmt.Printf("%s : %s\n", k, v)
	// }

	for v := range emails {
		fmt.Println("email address: " + v)
	}

	//pointers
	
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)
	fmt.Println(*&a)

	*b = 10
	fmt.Println(*b)
}
