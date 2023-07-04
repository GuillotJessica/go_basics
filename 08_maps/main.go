package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	//  Assign key/value
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Mike"] = "Mike@gmail.com"

	//  Declare map ens add key/value
	emails := map[string]string{"Bob":"bob@gmail.com", "Sharon":"sharon@gmail.com"}
	emails["Mike"] = "Mike@gmail.com"


	fmt.Println((emails))
	fmt.Println(len(emails))
	fmt.Println((emails["Bob"]))

	//  delete from map
	delete(emails, "Bob")
	fmt.Println((emails))
	fmt.Println(len(emails))
}