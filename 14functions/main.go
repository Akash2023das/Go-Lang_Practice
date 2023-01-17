package main

import "fmt"

//function to add two numbers
func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("functions in go")

	//taking two number and pass to add function
	sum := add(10, 20)
	fmt.Println("sum = ", sum)
}
