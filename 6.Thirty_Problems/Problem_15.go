/*
Write a function that takes a slice of integers and returns true if the slice is sorted in ascending order.
Input: [1, 2, 3, 4, 5]
Output: true
*/

package main

import "fmt"

//creating a function that returns true if the slice sorted 
func isSorted(numbers []int) bool {
	//storing the length of the slice
	n := len(numbers)

	//iterating over the slice numbers
	for i := 1; i < n; i++ { 
		//checking the condition pointing the second and first element
		if numbers[i] < numbers[i-1] { 
			return false
		}

	}
	return true
}
func main() {
	//slice of integers
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(isSorted(numbers))
}
