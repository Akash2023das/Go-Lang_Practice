package main

import (
	"fmt"
	"sort"
)

func main() {
	//slice in go

	//declare and initialize a string slice
	var colors = []string{"Red", "Blue", "Green"}

	//add a new color to the slice
	colors = append(colors, "Purple", "Pink", "Yellow")

	//printting the slice
	fmt.Println("Colors:", colors)

	//slicing the slice
	fmt.Println("Colors[1:4]:", colors[1:4]) //this will print from index 1 to 3

	//making an empty slice of size 5
	numbers := make([]int, 5)

	//adding values to the slice
	numbers[0] = 200
	numbers[1] = 500
	numbers[2] = 400
	numbers[3] = 300
	numbers[4] = 500
	//numbers[5] = 600 //this will throw an error

	//append a new value to the slice
	numbers = append(numbers, 600, 900, 100) //reinitializing the slice and adding new values

	//printting the slice
	fmt.Println("Numbers:", numbers)

	//sort the slice
	sort.Ints(numbers) //this will sort the slice in ascending order
	fmt.Println("Sorted Numbers:", numbers)

	//declare adn initialize a slice of strings
	var names = []string{"Javascript", "Python", "Go", "Java", "C++", "C", "Ruby", "PHP"}

	//remove  slice element based on index
	names = append(names[:2], names[3:]...) //this will remove the element at index 2

	fmt.Println("Names:", names)
}
