/*
Write a function that takes a slice of integers and returns the smallest number that is divisible by all elements in the slice.
Input: [2, 3, 5]
Output: 30
*/

package main

import "fmt"

func smallestDiv(numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		if i%numbers[i]==0
	}
}
func main() {
	numbers := []int{2, 3, 5}
	fmt.Println(smallestDiv(numbers))
}
