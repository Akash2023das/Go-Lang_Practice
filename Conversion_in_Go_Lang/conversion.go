package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our application.")
	fmt.Println("Please rate us between 1 to 5: ")

	reader := bufio.NewReader(os.Stdin) //reading input from keyboard

	input, _ := reader.ReadString('\n') //taking input as a string
	fmt.Println("Thanks for rating: ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //converting string to number and also trim the string for('\n)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}

}
