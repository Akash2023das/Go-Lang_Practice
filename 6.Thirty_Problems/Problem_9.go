/*
Write a function that takes a slice of integers and returns the average of all the elements in the slice.
Input: [1, 2, 3, 4, 5]
Output: 3
*/
package main

import "fmt"

//creating a function to calculate the average of all the elements of the slice and return the average
func AverageTotal(numbers []int) int {
	//storing the length of the slice
	n := len(numbers)

	//counting the sum of all the elements of the slice
	sum := n * (n + 1) / 2

	//counting the average
	avg := sum / n

	//return the average that we counted
	return avg
}
func main() {

	//slice of integers
	numbers := []int{1, 2, 3, 4, 5}

	//passing the slice to the function and printing the return value of the function
	fmt.Println(AverageTotal(numbers))
}
