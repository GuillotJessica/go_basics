package main

import (
	"fmt"
	"strconv"
)

//  Define person struct
type Person struct {
	// firstName string
	// lastName string
	// city string
	// gender string
	// age int

	firstName ,	lastName ,	city ,	gender string 
	age int
}

//   Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " is " + strconv.Itoa(p.age)
}

// has Birthday method (pointer receiver)
// has Birthday method (pointer receiver
func (p *Person) hasBirthday() {
	p.age++
}
	//  getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
	return
	}else {
		p.lastName = spouseLastName
	}
}
	
func main() {
	//  Init person using struc
	person1 := Person{firstName: "Lucy",
	 lastName: "Doigt",
	 city: "Paris",
	 gender: "F",
	 age: 35,
	}
	person2 := Person{firstName: "BoB",
	 lastName: "Wright",
	 city: "San Francisco",
	 gender: "M",
	 age: 38,
	}
//  fmt.Println(person1)
//  fmt.Println(person1.age)
//  person1.age ++
//  fmt.Println(person1.age)

// fmt.Println(person1.greet())
// person1.hasBirthday()
// fmt.Println(person1.greet())

fmt.Println(person1, person2)
person2.getMarried(person1.lastName)
person1.getMarried(person2.lastName)
fmt.Println(person1, person2)


}

