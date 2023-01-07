/*
Write a function that takes a slice of integers and returns a new slice with only the even numbers.
Input: [1, 2, 3, 4, 5]
Output: [2, 4]
*/
package main

import "fmt"

//creating a function that returns a slice of integers with even numbers
func evenNumbers(intSlice []int) []int {
	var ans[]int //taking an empty slice of integers

	//looping over the given slice of integers only taking the elements of the slice
	for _, i := range intSlice {

		//checking the condition ony by one 
		if i%2 == 0 {
			ans = append(ans, i) //putting the elements into the ans slice those who are even
		}
	}
	return ans //return the new ans slice that we created
}
func main() {

	//a slice of integers
	intSlice := []int{1, 2, 3, 4, 5}

	//passing the slice through the function and printing the result
	fmt.Println(evenNumbers(intSlice))
}
