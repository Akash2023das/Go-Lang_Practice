/*
Write a function that takes a slice of integers and returns the index of the first occurrence of a given number.
Input: ([1, 2, 3, 2, 4, 5], 2)
Output: 1
*/

package main

import "fmt"

//creating a function that returns index of the first occurrence of a given number
func indexOfElement(numbers []int, num int) int {

	//iterating over the numbers taking index and values
	for k, v := range numbers {

		//checking if the value is matched  then return the index
		if v == num {
			return k
		}
	}
	return -1 //if not found then return -1
}
func main() {
	//slice of integers
	numbers := []int{1, 2, 3, 2, 4, 5}

	//integer variable num
	num := 2

	//passing slice of integers and the given number into the function and printing the return value of the function
	fmt.Println(indexOfElement(numbers, num))
}
