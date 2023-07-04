package main

import "fmt"

func main(){
	var age  = 37
	const isCool bool = true

	// shorthand
	// name := "Jessica"
	// email := "blabla@gmail.com"
	name, email := "Jessica", "dsfdsf@gmail.com"

	fmt.Println(age, isCool, name, email)
	fmt.Printf("%T\n", name)
}
