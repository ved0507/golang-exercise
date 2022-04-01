package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to maths class of GOLANG")
	rand.Seed(time.Now().UnixMilli())
	time_now := time.Now().UnixMilli()
	fmt.Println(time_now)

	// var num1 int = 2
	// var num2 float64 = 4.456
	// fmt.Println(num1 + int(num2))
	fmt.Println(rand.Intn(100))

}
