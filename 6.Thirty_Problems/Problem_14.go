/*
Write a function that takes a slice of integers and returns a new slice with the elements in reverse order.
Input: [1, 2, 3, 4, 5]
Output: [5, 4, 3, 2, 1]
*/

package main

import "fmt"

//creating a function that returns a slice of integers
func reverseOrder(numbers []int) []int {
	//creating a empty slice of integers
	var ans []int

	//iterating the slice from the last position
	for i := len(numbers)-1; i >=0; i-- {
		ans = append(ans, numbers[i]) //adding the elements to the slice
	}
	return ans //returns the slice 
}
func main() {
	//slice of integers
	numbers := []int{10, 20, 30, 40, 50}

	fmt.Println(reverseOrder(numbers))
}
