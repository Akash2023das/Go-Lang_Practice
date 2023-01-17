package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go Lang")

	// var ptr *int
	// fmt.Println("Value of ptr is ", ptr)

	//taking a variable 10
	myNumber := 10

	//puting the address of myNumber in ptr
	ptr := &myNumber //storing the address of myNumber in ptr
	fmt.Println("Address of myNumber is ", &myNumber)
	fmt.Println("Value of ptr is ", ptr)  //this will print the address of myNumber
	fmt.Println("Value of ptr is ", *ptr) //this will print the value of myNumber

	//changing the value of myNumber
	*ptr = *ptr + 2
	fmt.Println("Value of myNumber is ", myNumber)
}
