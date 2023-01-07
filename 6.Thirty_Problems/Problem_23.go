/*
Write a function that takes a slice of integers and returns true if the slice contains a given number.
Input: ([1, 2, 3, 4, 5], 3)
Output: true
*/

package main

import "fmt"

func isPresent(numbers []int, num int) bool {
	for _, v := range numbers {
		if v == num {
			return true
		}
	}
	return false
}
func main() {
	numbers := []int{1, 2, 3, 4, 5}

	num := 3
	fmt.Println(isPresent(numbers, num))
}
