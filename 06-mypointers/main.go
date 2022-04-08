package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointers")

	// var ptr *int
	// another keyword for pointer ampersend sign --> &
	// fmt.Println("Value of Pointer is ", ptr)  // result -> nil

	// create my num
	myNumber := 25

	var ptr = &myNumber // reference means ampersend
	fmt.Println("Value of Pointer is ", ptr)
	fmt.Println("Value of Pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("new value is ", myNumber)

}
