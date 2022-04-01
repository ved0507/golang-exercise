package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza app")
	fmt.Println("Please rate our Pizza ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// message
	fmt.Println("Thanks for rating, ", input)

	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating:", numrating+1)
	}

}
