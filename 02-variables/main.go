package main

import (
	"fmt"
)

const loginToken string = "ghabbjhd" //public

func main() {
	var username string = "Vedavyas"
	fmt.Println(username)
	fmt.Printf("Variable type : %T \n", username)

	var isloggedin bool = false
	fmt.Println(isloggedin)
	fmt.Printf("Variable type : %T \n", isloggedin)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("Variable type : %T \n", smallval)

	var smallflat float32 = 255.5677
	fmt.Println(smallflat)
	fmt.Printf("Variable type : %T \n", smallflat)

	// Default variables
	var smallanother string
	fmt.Println(smallanother)
	fmt.Printf("Variable type : %T \n", smallanother)

	// implicit type
	var website = "www.vedavyas.com"
	fmt.Println(website)

	// No var style
	numberofuser := 300000
	fmt.Println(numberofuser)
}
