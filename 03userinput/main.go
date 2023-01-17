package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the program! This will show you how to get user input in Go.")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter your name: ")

	//comma ok syntax || comma error syntax
	name, _ := reader.ReadString('\n') // _ is use for ignoring the error
	fmt.Println("Hello My name is : ", name)

	//printing the type of the input
	fmt.Printf("The type of the input is %T", name)
}
