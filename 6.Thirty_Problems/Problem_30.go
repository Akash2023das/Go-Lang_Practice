/*
Write a function that takes a slice of integers and returns true if the slice is sorted in descending order.
 Input: [5, 4, 3, 2, 1]
 Output: true
*/

package main

import "fmt"

func checkSorted(numbers []int) bool {
	n := len(numbers)

	//iterating over the slice numbers
	for i := 1; i < n; i++ {
		//checking the condition pointing the second and first element
		if numbers[i] > numbers[i-1] {
			return false
		}

	}
	return true
}

func main() {
	numbers := []int{5, 4, 3, 2, 1}
	fmt.Println(checkSorted(numbers))
}
