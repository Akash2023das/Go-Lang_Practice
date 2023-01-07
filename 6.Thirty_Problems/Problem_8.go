/*
Write a function that takes a slice of integers and returns a new slice with all duplicates removed.
Input: [1, 2, 3, 2, 4, 5, 5]
Output: [1, 3, 4]
*/

package main

import "fmt"

func RemoveDuplicates(numbers []int) []int {

	//taking an empty map
	mk := make(map[int]int)

	//iterating the slice and adding keys and values into the map
	for _, n := range numbers {
		val, ok := mk[n]
		if ok {
			mk[n] = val + 1
		} else {
			mk[n] = 1
		}
	}

	//taking a empty slice of integers
	var ans []int

	//iterating over the map key and values
	for k, v := range mk {

		//checking those value that have only 1 key
		if v == 1 {
			ans = append(ans, k) //adding the key to the slice
		}
	}
	return ans //return the slice
}
func main() {

	//slice of integers
	numbers := []int{1, 2, 3, 2, 4, 5, 5}

	//passing the slice of integers into the function ans printing the return value of the function
	fmt.Println(RemoveDuplicates(numbers))
}
