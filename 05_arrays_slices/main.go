package main

import "fmt"

func main(){
	// Arrays
	// var fruitArr [2]string

	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and assign
	// fruitArr := [2]string{"Apple", "Orange"}
	
	// fmt.Println(fruitArr)
	
	//  Slice : not define number of keys/values
	fruitSlice := []string{"Apple", "Orange", "Grape"}
	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[1:3])
	fmt.Println(len(fruitSlice))
}