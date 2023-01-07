/*
Write a function that takes a slice of integers and returns a new slice with the elements in alternating order.
Input: [1, 2, 3, 4, 5]
Output: [1, 3, 5, 2, 4]
*/

package main

import "fmt"

func alternateNum(numbers []int) []int {
	var arr1 []int
	var arr2 []int
	var arr3 []int
	for i, v := range numbers {
		if i%2 == 0 {
			arr1 = append(arr1, v)
		} else {
			arr2 = append(arr2, v)
		}
	}
	arr3 = append(arr1, arr2...)

	return arr3
}
func main() {
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(alternateNum(numbers))
}
