package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	// sam := Person{"Sam", "Hollyer", 39} -> Depends on the order of the fields and not stable if changed
	// sam := Person{firstName: "Sam", lastName: "Hollyer", age: 39}

	var sam Person // -> Zero value {"","", 0}

	fmt.Println(sam)
	fmt.Printf("%+v", sam) // -> {firstName: lastName: age:0}

	sam.firstName = "Sam"
	sam.lastName = "Hollyer"
	sam.age = 39

	fmt.Printf("%+v", sam) // -> {firstName: lastName: age:0}
}
