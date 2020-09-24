package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	age       int
	height    float64
	employee  bool
}

// Embedded Structs
type person2 struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Andersen",
		age:       29,
		height:    1.70,
		employee:  true,
	}
	fmt.Println(alex.firstName)
	fmt.Println("Is employee?", alex.employee)

	// Check full struct
	//fmt.Printf("%+v", alex)

	uli := person2{
		firstName: "Ulises",
		lastName:  "Vargas",
		contactInfo: contactInfo{
			email: "texanico@gmail.com",
			zip:   1060,
		},
	}

	// Invoke Receiver function
	uli.print()

}

// Receiver Functions with structs

func (p person2) print() {
	fmt.Println("Ulises Email", p.contactInfo.email)
}
