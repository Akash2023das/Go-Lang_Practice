package main

import "fmt"

func main() {
	fmt.Println("loops in go")

	//declare a slice of strings
	fruits := []string{"apple", "banana", "grapes", "mango"}

	//for loop
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	//for range loop
	for i, fruit := range fruits {
		fmt.Printf("fruits[%d] = %s\n", i, fruit)
	}

	//for range loop without index
	for _, fruit := range fruits {
		fmt.Printf("fruit = %s\n", fruit)
	}

	//while loop
	i := 0
	for i < len(fruits) {
		fmt.Println(fruits[i])
		i++
	}
}
