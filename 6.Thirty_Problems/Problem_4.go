/*
Write a function that takes a slice of integers and returns the sum of all the elements in the slice.
Input: [1, 2, 3, 4, 5]
Output: 15
*/

package main

import "fmt"

//creating a function that returns sum of all the elements in a slice of integers
func sumOfSliceElements(intSlice []int) int {

	//storing the length of the slice
	n := len(intSlice)

	//calculating the sum of all the elements
	ans := n * (n + 1) / 2

	//return the result that we have calculated
	return ans
}
func main() {

	//slice of integers
	intSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(sumOfSliceElements(intSlice))
}
