/*
Write a function that takes a slice of strings and a string, and returns true if the string is contained in the slice.
Input: (["cat", "dog", "elephant"], "dog")
Output: true
*/

package main

import "fmt"

//creating a function that returns a boolean value
func isPresent(str []string, s string) bool {

	//iterating over the slice
	for _, v := range str {

		//checking if the given string is matching or not
		if v == s {
			return true
		}
	}

	//returning the ans
	return false
}
func main() {
	//slice of strings
	str := []string{"cat", "dog", "elephant"}

	//string variable s
	s := "dog"

	//passing the slice and the string to the function and printing the return value of the function
	fmt.Println(isPresent(str, s))
}
