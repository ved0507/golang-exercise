package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")

	presentTime := time.Now()
	fmt.Println(presentTime)
	// Standdard format 01-02-2006 15:04:05 Monday
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// create date
	createDate := time.Date(2022, time.May, 20, 12, 20, 0, 0, time.UTC)
	fmt.Println(createDate)
	//format
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))
}
