/*
Write a function that takes a slice of integers and returns the largest number in the slice.
Input: [1, 2, 3, 4, 5]
Output: 5
*/

package main

import "fmt"

//creating a function that takes a slice of integers and returns the largest number in the slice
func largestNumber(intSlice []int) int {

	//taking a variable assume the large value is zero
	large := 0

	//iterating over the slice of integers
	for i := 0; i < len(intSlice); i++ {

		//checking if the intSlice[i] is greater than the variable large
		if intSlice[i] > large {
			large = intSlice[i] //storing the largest value
		}
	}
	return large //return the largest value

}
func main() {
	//slice of integers
	intSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(largestNumber(intSlice))
}
